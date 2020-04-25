package types

import "io"

// publishes events to the event pipe
type Publisher interface {
	io.Closer
	Publish(event Event) error
}
