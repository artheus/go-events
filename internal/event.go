package internal

import (
	"github.com/artheus/go-events/types"
	"time"
)

type event struct {
	t time.Time
	o interface{}
}

func (e *event) Timestamp() *time.Time {
	return &e.t
}

func (e *event) Object() interface{} {
	return e.o
}

func Event(time time.Time, obj interface{}) types.Event {
	return &event{
		t: time,
		o: obj,
	}
}
