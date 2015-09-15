package logrus

import (
	"github.com/Sirupsen/logrus"
	"go.pedge.io/dlog"
)

func init() {
	dlog.SetLogger(logrus.New())
}
