package main

import (
	"fleet_management/internal/api"
	client "fleet_management/internal/client/db"
	mqttclient "fleet_management/internal/client/mqtt"
	rabbitmqclient "fleet_management/internal/client/rabbitmq"
	"fleet_management/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	config.Load()

	app := fiber.New()

	db := client.InitPostgreSQL()
	rabbitMQClient, err := rabbitmqclient.NewRabbitMQClient(config.Cfg.RabbitMQURL)
	if err != nil {
		logrus.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	mqttClient := mqttclient.NewMQTTClient(config.Cfg.MQTTHost, config.Cfg.MQTTPort, "fleet-backend", db, rabbitMQClient)

	if err = mqttClient.Connect(); err != nil {
		logrus.Fatalf("Failed to connect to MQTT broker: %v", err)
	}

	api.SetupRoutes(app, db, mqttClient, rabbitMQClient)

	logrus.Info("Server started at :3000")
	if err := app.Listen(":3000"); err != nil {
		logrus.Fatal(err)
	}
}
