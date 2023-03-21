package pubsub

import (
	"time"

	"github.com/ochom/orctic-library/utils"
	"github.com/ochom/pubsub"
	"github.com/ochom/quickmq/clients"
)

// Pubsub is an interface for pubsub implementations
type Pubsub interface {
	// Publish publishes a message to a channel
	Publish(channel string, message []byte, delay time.Duration) error

	// Subscribe subscribes to a channel
	Subscribe(channel string, handler func([]byte)) error
}

// QuickMQ is a quickmq implementation of the Pubsub interface
type QuickMQ struct {
	url string
}

// NewQuickMQ returns a new QuickMQ instance
func NewQuickMQ() *QuickMQ {
	quickURL := utils.GetEnv("QUICK_URL", "ws://localhost:3456")
	return &QuickMQ{
		url: quickURL,
	}
}

// Publish publishes a message to a channel
func (q *QuickMQ) Publish(channel string, message []byte, delay time.Duration) error {
	p := clients.NewPublisher(q.url, channel)
	if delay == 0 {
		return p.Publish(message)
	}

	return p.PublishWithDelay(message, delay)
}

// Subscribe subscribes to a channel
func (q *QuickMQ) Subscribe(channel string, handler func([]byte)) error {
	c := clients.NewConsumer(q.url, channel)
	return c.Consume(handler)
}

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
