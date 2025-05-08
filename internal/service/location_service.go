package service

import (
	"context"
	"encoding/json"
	client "fleet_management/internal/client/rabbitmq"
	"fleet_management/internal/config"
	"fleet_management/internal/dto"
	"fleet_management/internal/model"
	"fleet_management/internal/repository"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type LocationService interface {
	GetLatestLocation(ctx context.Context, vehicleID string) (*model.VehicleLocation, error)
	GetHistory(ctx context.Context, vehicleID string, start, end int64) ([]model.VehicleLocation, error)
	ProcessLocation(location dto.LocationRequest) error
}

type locationService struct {
	repo           repository.LocationRepository
	validator      *validator.Validate
	rabbitmqClient *client.RabbitMQClient
}

func NewLocationService(repo repository.LocationRepository, rabbitmqClient *client.RabbitMQClient) LocationService {
	return &locationService{
		repo:           repo,
		validator:      validator.New(),
		rabbitmqClient: rabbitmqClient,
	}
}

func (s *locationService) GetLatestLocation(ctx context.Context, vehicleID string) (*model.VehicleLocation, error) {
	return s.repo.GetLatestLocation(ctx, vehicleID)
}

func (s *locationService) GetHistory(ctx context.Context, vehicleID string, start, end int64) ([]model.VehicleLocation, error) {
	return s.repo.GetHistory(ctx, vehicleID, start, end)
}

func (s *locationService) ProcessLocation(location dto.LocationRequest) error {
	if location.VehicleID == "" || location.Latitude == 0 || location.Longitude == 0 {
		logrus.Errorf("Invalid location data: %+v", location)
		return fmt.Errorf("invalid location data")
	}

	locationModel := model.VehicleLocation{
		VehicleID: location.VehicleID,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		Timestamp: location.Timestamp,
	}

	err := s.repo.SaveLocation(locationModel)
	if err != nil {
		logrus.Errorf("Failed to save location: %v", err)
		return err
	}

	errGroup, _ := errgroup.WithContext(context.Background())
	errGroup.Go(func() error {
		return s.ValidateLocationRadius(location)
	})

	if err := errGroup.Wait(); err != nil {
		logrus.Errorf("Failed to validate location radius: %v", err)
	}

	logrus.Infof("Location data saved for vehicle %s", location.VehicleID)
	return nil
}

func (s *locationService) ValidateLocationRadius(location dto.LocationRequest) error {
	// Mock data for geofence, representing a bus station
	geofenceLat := config.Cfg.GeofenceLat
	geofenceLon := config.Cfg.GeofenceLon
	geofenceRadius := config.Cfg.GeofenceRadiusM

	distance := haversine(location.Latitude, location.Longitude, geofenceLat, geofenceLon)

	if distance <= geofenceRadius {
		log.Printf("Vehicle %s entered geofence, sending event", location.VehicleID)

		event := map[string]interface{}{
			"vehicle_id": location.VehicleID,
			"event":      "geofence_entry",
			"location": map[string]float64{
				"latitude":  location.Latitude,
				"longitude": location.Longitude,
			},
			"timestamp": time.Now().Unix(),
		}

		jsonData, _ := json.Marshal(event)
		if err := s.rabbitmqClient.PublishGeofenceEvent(jsonData); err != nil {
			log.Printf("Failed to publish geofence event: %v", err)
		}
	}

	return nil
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*math.Sin(dLon/2)*math.Sin(dLon/2)
	return R * 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}
