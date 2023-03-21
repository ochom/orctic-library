package pubsub

import (
	"time"

	"github.com/ochom/orctic-library/utils"
	"github.com/ochom/quickmq/clients"
)

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
