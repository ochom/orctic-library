package queues

import (
	"time"

	"github.com/ochom/orctic-library/utils"
	"github.com/ochom/pubsub"
)

// RabbitMQ is a rabbitmq implementation of the Pubsub interface
type RabbitMQ struct {
	url string
}

// NewRabbitMQ returns a new RabbitMQ instance
func NewRabbitMQ() *RabbitMQ {
	rabbitURL := utils.GetEnv("RABBIT_URL", "amqp://guest:guest@localhost:5672/")
	return &RabbitMQ{
		url: rabbitURL,
	}
}

// Publish publishes a message to a channel
func (r *RabbitMQ) Publish(channel string, message []byte, delay time.Duration) error {
	p := pubsub.NewPublisher(r.url, channel)
	if delay == 0 {

		return p.Publish(message)
	}

	return p.PublishWithDelay(message, delay)
}

// Subscribe subscribes to a channel
func (r *RabbitMQ) Subscribe(channel string, handler func([]byte)) error {
	c := pubsub.NewConsumer(r.url, channel)
	return c.Consume(handler)
}
