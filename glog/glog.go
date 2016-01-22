package glog // import "go.pedge.io/dlog/glog"

import (
	"go.pedge.io/dlog"

	"github.com/golang/glog"
)

func init() {
	dlog.SetLogger(NewLogger())
}

// NewLogger returns a new dlog.Logger for glog.
func NewLogger() dlog.Logger {
	return dlog.NewLogger(
		glog.Infoln,
		map[dlog.Level]func(...interface{}){
			dlog.LevelDebug: glog.Infoln,
			dlog.LevelInfo:  glog.Infoln,
			dlog.LevelWarn:  glog.Warningln,
			dlog.LevelError: glog.Errorln,
			// this relies on implementation-specific details of the generic dlog Logger implementation, not good
			dlog.LevelFatal: glog.Fatalln,
			dlog.LevelPanic: glog.Errorln,
		},
	)
}
