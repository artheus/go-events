package events

import (
	"github.com/artheus/go-events/internal"
	"github.com/sirupsen/logrus"
)

// set internal logger instance
func SetLogger(l logrus.FieldLogger) {
	internal.Logger = l
}

// set log level of internal logger
func SetLogLevel(lvl logrus.Level) {
	if l, ok := internal.Logger.(*logrus.Logger); !ok {
		internal.Logger.Error("unable to set log level of internal logger")
	} else {
		l.SetLevel(lvl)
	}
}