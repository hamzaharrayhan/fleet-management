package service

import (
	"encoding/json"
	client "fleet_management/internal/client/rabbitmq"
	"fleet_management/internal/dto"
	"fmt"
	"log"
)

type GeofenceWorker struct {
	rabbitmqClient *client.RabbitMQClient
}

func NewGeofenceWorker(rmq *client.RabbitMQClient) *GeofenceWorker {
	return &GeofenceWorker{rabbitmqClient: rmq}
}

func (w *GeofenceWorker) StartConsuming() {
	msgs, err := w.rabbitmqClient.Channel.Consume(
		"geofence_alerts",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(fmt.Sprintf("[GEOFENCE WORKER] | failed to start consuming: %v", err))
	}

	for msg := range msgs {
		var event dto.GeofenceEvent
		if err := json.Unmarshal(msg.Body, &event); err != nil {
			fmt.Println("[GEOFENCE WORKER] | Failed to parse message:", err)
			log.Printf("[GEOFENCE WORKER] | Failed to parse message: %v", err)
			continue
		}
		fmt.Printf("[GEOFENCE WORKER] | Received geofence event: %+v\n", event)
		log.Printf("[GEOFENCE WORKER] | Received geofence event: %+v\n", event)
	}
}
