package controller

import (
	"encoding/json"
	"fleet_management/internal/dto"
	"fleet_management/internal/service"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

type LocationMQTTController struct {
	locationService service.LocationService
}

func NewLocationMQTTController(locationService service.LocationService) *LocationMQTTController {
	return &LocationMQTTController{
		locationService: locationService,
	}
}

func (c *LocationMQTTController) HandleLocation(client mqtt.Client, msg mqtt.Message) {
	var location dto.LocationRequest
	if err := json.Unmarshal(msg.Payload(), &location); err != nil {
		logrus.Errorf("Failed to unmarshal message: %v", err)
		return
	}

	err := c.locationService.ProcessLocation(location)
	if err != nil {
		logrus.Errorf("Failed to process location for vehicle %s: %v", location.VehicleID, err)
		return
	}

	logrus.Infof("Successfully processed location for vehicle %s", location.VehicleID)
}
