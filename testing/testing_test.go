package testing

import (
	"log"
	"os"
	"testing"

	"github.com/Sirupsen/logrus"

	"go.pedge.io/dlog"
	dlogglog "go.pedge.io/dlog/glog"
	dloglion "go.pedge.io/dlog/lion"
	dloglogrus "go.pedge.io/dlog/logrus"
	dlogprotolog "go.pedge.io/dlog/protolog"
	"go.pedge.io/lion"
	"go.pedge.io/protolog"
)

func TestPrint(t *testing.T) {
	testPrint(t, dlog.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags)))
}

func TestPrintGlog(t *testing.T) {
	testPrint(t, dlogglog.NewLogger())
}

func TestPrintLion(t *testing.T) {
	testPrint(t, dloglion.NewLogger(lion.GlobalLogger()))
}

func TestPrintLogrus(t *testing.T) {
	testPrint(t, dloglogrus.NewLogger(logrus.StandardLogger()))
}

func TestPrintProtolog(t *testing.T) {
	testPrint(t, dlogprotolog.NewLogger(protolog.GlobalLogger()))
}

func testPrint(t *testing.T, logger dlog.Logger) {
	logger.WithField("key", "value").WithField("int", 1).Infof("number %d", 2)
	logger.Warnln("warning line")
}
