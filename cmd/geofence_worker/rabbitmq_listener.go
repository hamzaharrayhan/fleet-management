package main

import (
	"log"
	"time"

	client "fleet_management/internal/client/rabbitmq"
	"fleet_management/internal/config"
	"fleet_management/internal/service"
)

func main() {
	config.Load()

	log.Println("[GEOFENCE WORKER] | Starting geofence worker...")

	var rmqClient *client.RabbitMQClient
	var err error

	maxRetries := 10
	retryDelay := 3 * time.Second

	for i := 1; i <= maxRetries; i++ {
		log.Printf("[GEOFENCE WORKER] | Attempting to connect to RabbitMQ (attempt %d/%d)...", i, maxRetries)
		rmqClient, err = client.NewRabbitMQClient(config.Cfg.RabbitMQURL)
		if err == nil {
			log.Println("[GEOFENCE WORKER] | Successfully connected to RabbitMQ.")
			break
		}

		log.Printf("[GEOFENCE WORKER] | Failed to connect to RabbitMQ: %v", err)
		time.Sleep(retryDelay)
	}

	if rmqClient == nil {
		log.Fatalf("[GEOFENCE WORKER] | Could not connect to RabbitMQ after %d attempts. Exiting...", maxRetries)
	}

	defer rmqClient.Conn.Close()
	defer rmqClient.Channel.Close()

	worker := service.NewGeofenceWorker(rmqClient)
	log.Println("[GEOFENCE WORKER] | Geofence worker started and waiting for messages...")
	worker.StartConsuming()
}
