package protolog // import "go.pedge.io/dlog/protolog"

import (
	"go.pedge.io/dlog"
	"go.pedge.io/protolog"
)

func init() {
	protolog.AddGlobalHook(
		func(protologLogger protolog.Logger) {
			dlog.SetLogger(newLogger(protologLogger))
		},
	)
}

type logger struct {
	dlog.BaseLogger
	l protolog.Logger
}

func newLogger(l protolog.Logger) *logger {
	return &logger{l, l}
}

func (l *logger) AtLevel(level dlog.Level) dlog.Logger {
	// TODO(pedge): don't be lazy, make an actual map between dlog.Level and protolog.Level
	// you just copy/pasted lion_level.go to dlog_level.go
	return newLogger(l.l.AtLevel(protolog.Level(level)))
}

func (l *logger) WithField(key string, value interface{}) dlog.Logger {
	return newLogger(l.l.WithField(key, value))
}

func (l *logger) WithFields(fields map[string]interface{}) dlog.Logger {
	return newLogger(l.l.WithFields(fields))
}
