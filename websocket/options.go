package websocket

import "errors"

type OptionHandler func(*config) error

// WithMessageQueueSize Customize message channel queue size
func WithMessageQueueSize(size int) OptionHandler {
	return func(c *config) error {
		if size < 0 {
			return errors.New("invalid queue size (should not less than 0)")
		}
		c.messageQueueSize = size
		return nil
	}
}
