package protolog // import "go.pedge.io/dlog/protolog"

import (
	"go.pedge.io/dlog"
	"go.pedge.io/protolog"
)

func init() {
	dlog.SetLogger(newLogger(protolog.DefaultLogger))
}

type logger struct {
	dlog.BaseLogger
	l protolog.Logger
}

func newLogger(l protolog.Logger) *logger {
	return &logger{l, l}
}

func (l *logger) WithField(key string, value interface{}) dlog.Logger {
	return newLogger(l.l.WithField(key, value))
}

func (l *logger) WithFields(fields map[string]interface{}) dlog.Logger {
	return newLogger(l.l.WithFields(fields))
}
