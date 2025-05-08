package main

import (
	"fleet_management/internal/config"
	"fleet_management/internal/service"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var vehicleIDs = []string{
	"B1001TJ", "B1002TJ", "B1003TJ", "B1004TJ", "B1005TJ",
}

func main() {
	config.Load()

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", config.Cfg.MQTTHost, config.Cfg.MQTTPort))
	opts.SetClientID("mqtt-client")

	client := mqtt.NewClient(opts)
	locationPublisherService := service.NewLocationPublisher(client)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("[MQTT PUBLISHER] | Failed to connect to MQTT broker: %v || %s:%s", token.Error(), config.Cfg.MQTTHost, config.Cfg.MQTTPort)
	}
	defer client.Disconnect(250)

	go locationPublisherService.PublishLocation(vehicleIDs)

	select {}
}
