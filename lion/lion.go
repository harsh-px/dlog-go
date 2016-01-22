package lion // import "go.pedge.io/dlog/lion"

import (
	"go.pedge.io/dlog"
	"go.pedge.io/lion"
)

func init() {
	dlog.SetLogger(newLogger(lion.DefaultLogger))
}

type logger struct {
	dlog.BaseLogger
	l lion.Logger
}

func newLogger(l lion.Logger) *logger {
	return &logger{l, l}
}

func (l *logger) WithField(key string, value interface{}) dlog.Logger {
	return newLogger(l.l.WithField(key, value))
}

func (l *logger) WithFields(fields map[string]interface{}) dlog.Logger {
	return newLogger(l.l.WithFields(fields))
}
