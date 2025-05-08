package model

type VehicleLocation struct {
	VehicleID string  `db:"vehicle_id"`
	Latitude  float64 `db:"latitude"`
	Longitude float64 `db:"longitude"`
	Timestamp int64   `db:"timestamp"`
}
