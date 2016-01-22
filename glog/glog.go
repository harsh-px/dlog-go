package glog // import "go.pedge.io/dlog/glog"

import (
	"fmt"

	"go.pedge.io/dlog"

	"github.com/golang/glog"
)

func init() {
	dlog.SetLogger(newLogger())
}

type logger struct{}

func newLogger() *logger {
	return &logger{}
}

func (l *logger) Debugf(format string, args ...interface{}) {
	glog.V(1).Infof(format, args...)
}

func (l *logger) Debugln(args ...interface{}) {
	glog.V(1).Infoln(args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	glog.Infof(format, args...)
}

func (l *logger) Infoln(args ...interface{}) {
	glog.Infoln(args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	glog.Warningf(format, args...)
}

func (l *logger) Warnln(args ...interface{}) {
	glog.Warningln(args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	glog.Errorf(format, args...)
}

func (l *logger) Errorln(args ...interface{}) {
	glog.Errorln(args...)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	glog.Fatalf(format, args...)
}

func (l *logger) Fatalln(args ...interface{}) {
	glog.Fatalln(args...)
}

func (l *logger) Panicf(format string, args ...interface{}) {
	glog.Infof(format, args...)
	panic(fmt.Sprintf(format, args...))
}

func (l *logger) Panicln(args ...interface{}) {
	glog.Infoln(args...)
	panic(fmt.Sprintln(args...))
}

func (l *logger) Printf(format string, args ...interface{}) {
	glog.Infof(format, args...)
}

func (l *logger) Println(args ...interface{}) {
	glog.Infoln(args...)
}
