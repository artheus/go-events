package internal

import (
	"errors"
	"github.com/artheus/go-events/types"
)

type subscriber struct {
	c chan types.Event
	closed bool
}

func (s *subscriber) Close() error {
	s.closed = true
	close(s.c)
	return nil
}

const (
	eventNotSentError = "event was not sent to subscriber"
	alreadyClosedError = "subscriber is already closed"
)

func (s *subscriber) send(e types.Event) error {
	if s.closed {
		return errors.New(alreadyClosedError)
	}

	select {
	case s.c <- e:
		break
	default:
		return errors.New(eventNotSentError)
	}

	return nil
}

func (s *subscriber) Get() <-chan types.Event {
	return s.c
}
