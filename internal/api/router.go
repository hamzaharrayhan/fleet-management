package api

import (
	mqttclient "fleet_management/internal/client/mqtt"
	rabbitmqclient "fleet_management/internal/client/rabbitmq"
	controller "fleet_management/internal/controller/http"
	"fleet_management/internal/repository"
	"fleet_management/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(app *fiber.App, db *sqlx.DB, mqttClient *mqttclient.MQTTClient, rabbitMQClient *rabbitmqclient.RabbitMQClient) {

	locationRepo := repository.NewLocationRepository(db)

	locationService := service.NewLocationService(locationRepo, rabbitMQClient)

	locationController := controller.NewLocationController(locationService)

	vehicle := app.Group("/vehicles")
	vehicle.Get("/:vehicle_id/location", locationController.GetLatestLocation)
	vehicle.Get("/:vehicle_id/history", locationController.GetHistory)

	go mqttClient.StartListening()
	app.Listen(":3000")
}
