package internal

import (
	"errors"
	"github.com/artheus/go-events/types"
)

type pipe struct {
	c chan types.Event
}

func NewPipe(chanSize int) types.Pipe {
	return &pipe{
		c: make(chan types.Event, chanSize),
	}
}

func (p *pipe) Close() error {
	if _, ok := <-p.c; !ok {
		return errors.New("pipe already closed")
	}

	close(p.c)
	return nil
}

func (p *pipe) Publisher() types.Publisher {
	return &publisher{c: p.c}
}

func (p *pipe) Subscriber() types.Subscriber {
	return &subscriber{c: p.c}
}
