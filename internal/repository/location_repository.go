package repository

import (
	"context"
	"fleet_management/internal/model"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type LocationRepository interface {
	SaveLocation(location model.VehicleLocation) error
	GetLatestLocation(ctx context.Context, vehicleID string) (*model.VehicleLocation, error)
	GetHistory(ctx context.Context, vehicleID string, start, end int64) ([]model.VehicleLocation, error)
}

type locationRepo struct {
	db *sqlx.DB
}

func NewLocationRepository(db *sqlx.DB) LocationRepository {
	return &locationRepo{db: db}
}

func (r *locationRepo) SaveLocation(location model.VehicleLocation) error {
	query := `INSERT INTO vehicle_locations (vehicle_id, latitude, longitude, timestamp) 
			  VALUES ($1, $2, $3, $4)`

	_, err := r.db.Exec(query, location.VehicleID, location.Latitude, location.Longitude, location.Timestamp)
	if err != nil {
		return fmt.Errorf("failed to insert vehicle location: %v", err)
	}
	return nil
}

func (r *locationRepo) GetLatestLocation(ctx context.Context, vehicleID string) (*model.VehicleLocation, error) {
	var loc model.VehicleLocation
	query := `SELECT vehicle_id, latitude, longitude, timestamp 
			  FROM vehicle_locations 
			  WHERE vehicle_id = $1 
			  ORDER BY timestamp DESC 
			  LIMIT 1`
	err := r.db.GetContext(ctx, &loc, query, vehicleID)
	if err != nil {
		logrus.WithError(err).Error("failed to get latest location")
		return nil, err
	}
	return &loc, nil
}

func (r *locationRepo) GetHistory(ctx context.Context, vehicleID string, start, end int64) ([]model.VehicleLocation, error) {
	var locations []model.VehicleLocation
	query := `SELECT vehicle_id, latitude, longitude, timestamp 
			  FROM vehicle_locations 
			  WHERE vehicle_id = $1 AND timestamp BETWEEN $2 AND $3 
			  ORDER BY timestamp ASC`
	err := r.db.SelectContext(ctx, &locations, query, vehicleID, start, end)
	if err != nil {
		logrus.WithError(err).Error("failed to get location history")
		return nil, err
	}
	return locations, nil
}
