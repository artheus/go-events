package internal

import "github.com/artheus/go-events/types"

type subscriber struct {
	c <-chan types.Event
}

func (s *subscriber) Close() error {
	return nil
}

func (s *subscriber) Get() <-chan types.Event {
	return s.c
}
