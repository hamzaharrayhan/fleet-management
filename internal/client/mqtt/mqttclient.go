package mqttclient

import (
	rabbitmqclient "fleet_management/internal/client/rabbitmq"
	"fleet_management/internal/controller"
	"fleet_management/internal/repository"
	"fleet_management/internal/service"
	"fmt"

	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type MQTTClient struct {
	client     mqtt.Client
	controller *controller.LocationMQTTController
}

func NewMQTTClient(host, port, clientID string, db *sqlx.DB, rabbitMQClient *rabbitmqclient.RabbitMQClient) *MQTTClient {
	logrus.Infof("Connecting to MQTT broker: %s:%s", host, port)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", host, port))
	opts.SetClientID(clientID)
	opts.SetAutoReconnect(true)
	opts.SetCleanSession(true)

	client := mqtt.NewClient(opts)

	repository := repository.NewLocationRepository(db)
	service := service.NewLocationService(repository, rabbitMQClient)
	controller := controller.NewLocationMQTTController(service)
	return &MQTTClient{client: client, controller: controller}
}

func (m *MQTTClient) Connect() error {
	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (m *MQTTClient) Listen(topic string, callback mqtt.MessageHandler) {
	if token := m.client.Subscribe(topic, 0, callback); token.Wait() && token.Error() != nil {
		log.Fatalf("Error subscribing to topic %s: %v", topic, token.Error())
	}
}
