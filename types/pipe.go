package types

import "io"

// event pipe holding the main channel
// instances of Publisher and Subscriber should only be provided from Pipe
type Pipe interface {
	io.Closer
	Publisher() Publisher
	Subscriber() Subscriber
}