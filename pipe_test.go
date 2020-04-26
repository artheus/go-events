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

func (s *PipeTestSuite) TearDownTest() {
	var err error

	err = s.p.Close()
	s.NoError(err)
}

func (s *PipeTestSuite) TestPipeShouldErrorOnChanSizeLessThanOne() {
	_, err := Pipe(0)
	s.Error(err)
}

func (s *PipeTestSuite) TestPubSub() {
	var err error
	var sub = s.p.Subscriber()
	var pub = s.p.Publisher()
	defer sub.Close()
	defer pub.Close()

	evtTime, evtObj := time.Now(), "test object"

	var evt = Event(evtTime, evtObj)

	err = pub.Publish(evt)
	s.NoError(err)

	eo := <-sub.Get()
	s.NotNil(eo)
	s.Equal(evtObj, eo.Object())
	s.Equal(evtTime, *eo.Timestamp())
}

func (s *PipeTestSuite) TestPubMultipleSub() {
	var err error
	var pub = s.p.Publisher()
	var subs = []types.Subscriber{
		s.p.Subscriber(),
		s.p.Subscriber(),
		s.p.Subscriber(),
		s.p.Subscriber(),
	}
	defer pub.Close()
	defer func() { // close all subscribers
		for _, sub := range subs {
			err = sub.Close()
			s.NoError(err)
		}
	}()

	evtTime, evtObj := time.Now(), "test object"

	var evt = Event(evtTime, evtObj)

	err = pub.Publish(evt)
	s.NoError(err)

	for _, sub := range subs {
		eo := <-sub.Get()
		s.NotNil(eo)
		s.Equal(evtObj, eo.Object())
		s.Equal(evtTime, *eo.Timestamp())
	}
}

func TestRunPipeTestSuite(t *testing.T) {
	suite.Run(t, new(PipeTestSuite))
}