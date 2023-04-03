package queues

import (
	"fmt"
	"time"

	"github.com/ochom/orctic-library/utils"
	"github.com/ochom/quickmq/clients"
)

var (
	quickURL = utils.GetEnv("QUICKMQ_URL", "ws://localhost:3456")
)

// quickMQ is a  quickmq implementation of the Queue interface
type quickMQ struct {
	url         string
	queuePrefix string
}

// NewQuickMQ returns a new  quickMQ instance
func NewQuickMQ() Queue {
	return &quickMQ{
		url:         quickURL,
		queuePrefix: queuePrefix,
	}
}

// Publish publishes a message to a channel
func (r *quickMQ) Publish(channel string, message []byte) error {
	channel = fmt.Sprintf("%s-%s", r.queuePrefix, channel)
	p := clients.NewPublisher(r.url, channel)
	return p.Publish(message)
}

// PublishWithDelay publishes a message to a channel with a delay
func (r *quickMQ) PublishWithDelay(channel string, message []byte, delay time.Duration) error {
	channel = fmt.Sprintf("%s-%s", r.queuePrefix, channel)
	p := clients.NewPublisher(r.url, channel)
	return p.PublishWithDelay(message, delay)
}

// Subscribe subscribes to a channel
func (r *quickMQ) Subscribe(channel string, handler func([]byte)) error {
	channel = fmt.Sprintf("%s-%s", r.queuePrefix, channel)
	c := clients.NewConsumer(r.url, channel)
	return c.Consume(handler)
}
