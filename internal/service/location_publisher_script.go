package service

import (
	"fleet_management/internal/config"
	"fmt"
	"math/rand/v2"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type LocationPublisher struct {
	mqttClient mqtt.Client
}

func NewLocationPublisher(mqttClient mqtt.Client) *LocationPublisher {
	return &LocationPublisher{mqttClient: mqttClient}
}

var vehicleIDs = []string{
	"B1001TJ", "B1002TJ", "B1003TJ", "B1004TJ", "B1005TJ",
}

func generateLocation() (float64, float64) {
	latitude := config.Cfg.GeofenceLat + (rand.Float64()-0.5)*0.001
	longitude := config.Cfg.GeofenceLon + (rand.Float64()-0.5)*0.001
	return latitude, longitude
}

func (l *LocationPublisher) PublishLocation(vehicleID []string) {
	for {
		for _, vehicleID := range vehicleIDs {
			latitude, longitude := generateLocation()
			payload := fmt.Sprintf(`{
				"vehicle_id": "%s",
				"latitude": %f,
				"longitude": %f,
				"timestamp": %d
			}`, vehicleID, latitude, longitude, time.Now().Unix())

			token := l.mqttClient.Publish(fmt.Sprintf("/fleet/vehicle/%s/location", vehicleID), 0, false, payload)
			token.Wait()

			time.Sleep(2 * time.Second)
		}
	}
}
