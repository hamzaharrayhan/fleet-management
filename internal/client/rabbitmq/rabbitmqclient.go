package client

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitMQClient struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQClient(rabbitmqUrl string) (*RabbitMQClient, error) {
	logrus.Infof("Connecting to RabbitMQ... URL: %s", rabbitmqUrl)
	conn, err := amqp.Dial(rabbitmqUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	err = ch.ExchangeDeclare(
		"fleet.events",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	_, err = ch.QueueDeclare(
		"geofence_alerts",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	err = ch.QueueBind(
		"geofence_alerts",
		"geofence_alerts",
		"fleet.events",
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind queue: %w", err)
	}

	return &RabbitMQClient{
		Conn:    conn,
		Channel: ch,
	}, nil
}
