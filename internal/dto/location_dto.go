package dto

type LocationResponse struct {
	VehicleID string  `json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp int64   `json:"timestamp"`
}

type HistoryQuery struct {
	Start int64 `query:"start"`
	End   int64 `query:"end"`
}

type LocationRequest struct {
	VehicleID string  `json:"vehicle_id" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
	Timestamp int64   `json:"timestamp" validate:"required"`
}
