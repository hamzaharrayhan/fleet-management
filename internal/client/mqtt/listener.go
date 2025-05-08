package mqttclient

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

func (l *MQTTClient) StartListening() {
	topic := "/fleet/vehicle/+/location"

	l.Listen(topic, func(client mqtt.Client, msg mqtt.Message) {
		l.controller.HandleLocation(client, msg)
	})

	logrus.Infof("Listening on MQTT topic: %s", topic)
}
