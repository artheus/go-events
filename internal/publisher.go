package internal

import (
	"errors"
	"github.com/artheus/go-events/types"
)

const (
	chanClosedErr = "pipe channel is closed"
)

type publisher struct {
	c chan<- types.Event
}

func (p *publisher) Close() error {
	return nil
}

func (p *publisher) Publish(event types.Event) (err error) {
	select {
	case p.c <- event:
		break
	default:
		return errors.New(chanClosedErr)
	}

	return nil
}
