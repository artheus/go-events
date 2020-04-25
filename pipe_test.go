package events

import (
	"github.com/artheus/go-events/types"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

const testChanSize = 5

type PipeTestSuite struct {
	suite.Suite
	p types.Pipe
}

func (s *PipeTestSuite) SetupTest() {
	var err error

	s.p, err = Pipe(testChanSize)
	s.NoError(err)
}

func (s *PipeTestSuite) TestPubSub() {
	var err error
	var sub = s.p.Subscriber()
	var pub = s.p.Publisher()

	evtTime, evtObj := time.Now(), "test object"

	var evt = Event(evtTime, evtObj)

	err = pub.Publish(evt)
	s.NoError(err)

	eo := <-sub.Get()
	s.NotNil(eo)
	s.Equal(evtObj, eo.Object())
	s.Equal(evtTime, *eo.Timestamp())
}

func TestRunPipeTestSuite(t *testing.T) {
	suite.Run(t, new(PipeTestSuite))
}