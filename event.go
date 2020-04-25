package events

import (
	"github.com/artheus/go-events/internal"
	"github.com/artheus/go-events/types"
	"time"
)

// Create a new event type with provided timestamp and
// event object to provide to subscribers
func Event(t time.Time, o interface{}) types.Event {
	return internal.Event(t, o)
}