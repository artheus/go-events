package events

import (
	"errors"
	"github.com/artheus/go-events/internal"
	"github.com/artheus/go-events/types"
)

const tooSmallChannelSizeErr = "too small channel size"

// Creates a new event pipe with provided channel size
// returns error only if channel size is less than 1
func Pipe(channelSize int) (types.Pipe, error) {
	if channelSize < 1 {
		return nil, errors.New(tooSmallChannelSizeErr)
	}

	return internal.NewPipe(channelSize), nil
}