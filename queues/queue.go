package queues

import (
	"time"

	"github.com/ochom/orctic-library/utils"
)

// Queue ...
type Queue interface {
	// Publish publishes a message to a channel
	Publish(channel string, message []byte) error

	// PublishWithDelay publishes a message to a channel with a delay
	PublishWithDelay(channel string, message []byte, delay time.Duration) error

	// Subscribe subscribes to a channel
	Subscribe(channel string, handler func([]byte)) error
}

// queuePrefix is the prefix for all queues
var queuePrefix = utils.GetEnv("QUEUE_PREFIX", "dev")
