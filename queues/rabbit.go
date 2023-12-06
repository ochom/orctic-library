package queues

import (
	"fmt"
	"time"

	"github.com/ochom/gutils/helpers"
	"github.com/ochom/pubsub"
)

var (
	rabbitURL   = helpers.GetEnv("RABBIT_URL", "amqp://guest:guest@localhost:5672/")
	queuePrefix = helpers.GetEnv("QUEUE_PREFIX", "dev")
)

// Publish publishes a message to a channel
func Publish(channel string, message []byte, delay ...time.Duration) error {
	channel = fmt.Sprintf("%s-%s", queuePrefix, channel)
	p := pubsub.NewPublisher(rabbitURL, channel)
	return p.Publish(message)
}

// PublishWithDelay publishes a message to a channel with a delay
func PublishWithDelay(channel string, message []byte, delay time.Duration) error {
	channel = fmt.Sprintf("%s-%s", queuePrefix, channel)
	p := pubsub.NewPublisher(rabbitURL, channel)
	return p.PublishWithDelay(message, delay)
}

// Subscribe subscribes to a channel
func Subscribe(channel string, handler func([]byte)) error {
	channel = fmt.Sprintf("%s-%s", queuePrefix, channel)
	c := pubsub.NewConsumer(rabbitURL, channel)
	return c.Consume(handler)
}
