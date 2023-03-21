package pubsub

import (
	"time"

	"github.com/ochom/orctic-library/utils"
)

// Pubsub is an interface for pubsub implementations
type Pubsub interface {
	// Publish publishes a message to a channel
	Publish(channel string, message []byte, delay time.Duration) error

	// Subscribe subscribes to a channel
	Subscribe(channel string, handler func([]byte)) error
}

// New returns a new Pubsub instance
func New() Pubsub {
	provider := utils.GetEnv("PUBSUB_PROVIDER", "rabbitmq")
	if provider == "quickmq" {
		return NewRabbitMQ()
	}

	return NewRabbitMQ()
}
