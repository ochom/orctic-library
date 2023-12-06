package queues

import (
	"fmt"
	"time"

	"github.com/ochom/gutils/helpers"
	"github.com/ochom/pubsub"
)

var rabbitURL = helpers.GetEnv("RABBIT_URL", "amqp://guest:guest@localhost:5672/")

// rabbitMQ is a rabbitmq implementation of the Pubsub interface
type rabbitMQ struct {
	url         string
	queuePrefix string
}

// NewRabbitMQ returns a new RabbitMQ instance
func NewRabbitMQ() Queue {
	return &rabbitMQ{
		url:         rabbitURL,
		queuePrefix: queuePrefix,
	}
}

// Publish publishes a message to a channel
func (r *rabbitMQ) Publish(channel string, message []byte) error {
	channel = fmt.Sprintf("%s-%s", r.queuePrefix, channel)
	p := pubsub.NewPublisher(r.url, channel)
	return p.Publish(message)
}

// PublishWithDelay publishes a message to a channel with a delay
func (r *rabbitMQ) PublishWithDelay(channel string, message []byte, delay time.Duration) error {
	channel = fmt.Sprintf("%s-%s", r.queuePrefix, channel)
	p := pubsub.NewPublisher(r.url, channel)
	return p.PublishWithDelay(message, delay)
}

// Subscribe subscribes to a channel
func (r *rabbitMQ) Subscribe(channel string, handler func([]byte)) error {
	channel = fmt.Sprintf("%s-%s", r.queuePrefix, channel)
	c := pubsub.NewConsumer(r.url, channel)
	return c.Consume(handler)
}
