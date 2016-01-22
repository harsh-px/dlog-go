package logrus // import "go.pedge.io/dlog/logrus"

import (
	"go.pedge.io/dlog"

	"github.com/Sirupsen/logrus"
)

func init() {
	dlog.SetLogger(newLogger(logrus.New()))
}

type logrusLogger interface {
	dlog.BaseLogger
	WithField(key string, value interface{}) *logrus.Entry
	WithFields(fields logrus.Fields) *logrus.Entry
}

type logger struct {
	dlog.BaseLogger
	l logrusLogger
}

func newLogger(l logrusLogger) *logger {
	return &logger{l, l}
}

func (l *logger) WithField(key string, value interface{}) dlog.Logger {
	return newLogger(l.l.WithField(key, value))
}

func (l *logger) WithFields(fields map[string]interface{}) dlog.Logger {
	return newLogger(l.l.WithFields(fields))
}
