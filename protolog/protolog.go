package protolog // import "go.pedge.io/dlog/protolog"

import (
	"os"

	"go.pedge.io/dlog"
	"go.pedge.io/protolog"
)

func init() {
	dlog.SetLogger(NewLogger(protolog.NewStandardLogger(protolog.NewFileFlusher(os.Stderr))))
}

type logger struct {
	protolog.Logger
}

// NewLogger creates a new dlog.Logger using a protolog.Logger.
func NewLogger(l protolog.Logger) dlog.Logger {
	return &logger{l}
}

func (l *logger) Debug(args ...interface{}) {
	l.Debugln(args...)
}

func (l *logger) Info(args ...interface{}) {
	l.Infoln(args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.Warnln(args...)
}

func (l *logger) Error(args ...interface{}) {
	l.Errorln(args...)
}

func (l *logger) Fatal(args ...interface{}) {
	l.Fatalln(args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.Panicln(args...)
}

func (l *logger) Print(args ...interface{}) {
	l.Println(args...)
}
