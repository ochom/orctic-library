package queues

import (
	"fmt"
	"time"

	"github.com/ochom/orctic-library/utils"
	"github.com/ochom/pubsub"
)

// RabbitMQ is a rabbitmq implementation of the Pubsub interface
type RabbitMQ struct {
	url    string
	prefix string
}

// NewRabbitMQ returns a new RabbitMQ instance
func NewRabbitMQ() *RabbitMQ {
	rabbitURL := utils.GetEnv("RABBIT_URL", "amqp://guest:guest@localhost:5672/")
	var queuePrefix = utils.GetEnv("QUEUE_PREFIX", "dev")
	return &RabbitMQ{
		url:    rabbitURL,
		prefix: queuePrefix,
	}
}

// Publish publishes a message to a channel
func (r *RabbitMQ) Publish(channel string, message []byte) error {
	channel = fmt.Sprintf("%s-%s", r.prefix, channel)
	p := pubsub.NewPublisher(r.url, channel)
	return p.Publish(message)
}

// PublishWithDelay publishes a message to a channel with a delay
func (r *RabbitMQ) PublishWithDelay(channel string, message []byte, delay time.Duration) error {
	channel = fmt.Sprintf("%s-%s", r.prefix, channel)
	p := pubsub.NewPublisher(r.url, channel)
	return p.PublishWithDelay(message, delay)
}

// Subscribe subscribes to a channel
func (r *RabbitMQ) Subscribe(channel string, handler func([]byte), lazy bool) error {
	channel = fmt.Sprintf("%s-%s", r.prefix, channel)
	c := pubsub.NewConsumer(r.url, channel)
	return c.Consume(handler, lazy)
}
