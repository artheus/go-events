package types

import "io"

// subscribes to events published to the event pipe
type Subscriber interface {
	io.Closer
	Get() <-chan Event
}
