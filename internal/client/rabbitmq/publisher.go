package client

import (
	"github.com/streadway/amqp"
)

func (r *RabbitMQClient) PublishGeofenceEvent(body []byte) error {
	return r.Channel.Publish(
		"fleet.events",
		"geofence_alerts",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
