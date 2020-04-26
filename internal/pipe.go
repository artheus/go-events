package internal

import (
	"context"
	"github.com/artheus/go-events/types"
	"sync"
)

type pipe struct {
	c      chan types.Event
	s      []*subscriber
	mx     *sync.Mutex
	cs     int
	ctx    context.Context
	cancel context.CancelFunc
}

func NewPipe(chanSize int) types.Pipe {
	var ctx, cancel = context.WithCancel(context.Background())

	p := &pipe{
		c:      make(chan types.Event, chanSize),
		cs:     chanSize,
		mx:     &sync.Mutex{},
		ctx:    ctx,
		cancel: cancel,
	}

	go func() { // routine for pipe event broadcasting
		var err error

		for {
			select {
			case <-p.ctx.Done():
				return
			case e := <-p.c:
				if e == nil { // channel closed
					break
				}

				p.mx.Lock()
				for _, s := range p.s {
					if err = s.send(e); err != nil {
						Logger.Errorf("error broadcasting event to subscriber: %+v", err)
					}
				}
				p.mx.Unlock()
			}
		}
	}()

	return p
}

func (p *pipe) Close() error {
	p.cancel()
	close(p.c)

	return nil
}

func (p *pipe) Publisher() types.Publisher {
	return &publisher{c: p.c}
}

func (p *pipe) Subscriber() types.Subscriber {
	var s = &subscriber{c: make(chan types.Event, p.cs)}
	p.mx.Lock()
	p.s = append(p.s, s)
	p.mx.Unlock()
	return s
}
