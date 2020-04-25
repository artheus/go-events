package types

import "time"

// interface for event with timestamp and arbitrary object
type Event interface {
	Timestamp() *time.Time
	Object() interface{}
}