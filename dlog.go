/*
Package dlog (delegating log) wraps common functionality for common golang logging packages.

The Logger interface wraps the common logging functionality. Every method on Logger
is also a global method on the dlog package. Given an implementation of Logger, you can
register it as the global logger by calling:

	func register(logger dlog.Logger) {
	  dlog.SetLogger(logger)
	}

To make things simple, packages for glog, logrus, protolog, and lion are given with the ability to blank import:

	import (
	  _ "go.pedge.io/dlog/glog" // set glog as the global logger
	  _ "go.pedge.io/dlog/logrus" // set logrus as the global logger with default settings
	  _ "go.pedge.io/dlog/protolog" // set protolog as the global logger with default settings
	  _ "go.pedge.io/dlog/lion" // set lion as the global logger with default settings
	)

Or, do something more custom:

	func init() { // or anywhere
	  logger := logrus.New()
	  logger.Out = os.Stdout
	  logger.Formatter = &logrus.TextFormatter{
		ForceColors: true,
	  }
	  dlog.SetLogger(logger)
	}

By default, golang's standard logger is used.
*/
package dlog // import "go.pedge.io/dlog"

import (
	"log"
	"os"
)

var (
	globalLogger = NewLogger(log.New(os.Stderr, "", log.LstdFlags))
)

// Logger is an interface that all logging implementations must implement.
type Logger interface {
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
}

// SetLogger sets the global logger used by dlog.
func SetLogger(logger Logger) {
	globalLogger = logger
}

// NewLogger creates a new Logger using a standard log.Logger.
func NewLogger(l *log.Logger) Logger {
	return newLogger(l)
}

// Debugf logs at the debug level with the semantics of fmt.Printf.
func Debugf(format string, args ...interface{}) {
	globalLogger.Debugf(format, args...)
}

// Debugln logs at the debug level with the semantics of fmt.Println.
func Debugln(args ...interface{}) {
	globalLogger.Debugln(args...)
}

// Infof logs at the info level with the semantics of fmt.Printf.
func Infof(format string, args ...interface{}) {
	globalLogger.Infof(format, args...)
}

// Infoln logs at the info level with the semantics of fmt.Println.
func Infoln(args ...interface{}) {
	globalLogger.Infoln(args...)
}

// Warnf logs at the warn level with the semantics of fmt.Printf.
func Warnf(format string, args ...interface{}) {
	globalLogger.Warnf(format, args...)
}

// Warnln logs at the warn level with the semantics of fmt.Println.
func Warnln(args ...interface{}) {
	globalLogger.Warnln(args...)
}

// Errorf logs at the error level with the semantics of fmt.Printf.
func Errorf(format string, args ...interface{}) {
	globalLogger.Errorf(format, args...)
}

// Errorln logs at the error level with the semantics of fmt.Println.
func Errorln(args ...interface{}) {
	globalLogger.Errorln(args...)
}

// Fatalf logs at the fatal level with the semantics of fmt.Printf and exits with os.Exit(1).
func Fatalf(format string, args ...interface{}) {
	globalLogger.Fatalf(format, args...)
}

// Fatalln logs at the fatal level with the semantics of fmt.Println and exits with os.Exit(1).
func Fatalln(args ...interface{}) {
	globalLogger.Fatalln(args...)
}

// Panicf logs at the panic level with the semantics of fmt.Printf and panics.
func Panicf(format string, args ...interface{}) {
	globalLogger.Panicf(format, args...)
}

// Panicln logs at the panic level with the semantics of fmt.Println and panics.
func Panicln(args ...interface{}) {
	globalLogger.Panicln(args...)
}

// Printf logs at the info level with the semantics of fmt.Printf.
func Printf(format string, args ...interface{}) {
	globalLogger.Printf(format, args...)
}

// Println logs at the info level with the semantics of fmt.Println.
func Println(args ...interface{}) {
	globalLogger.Println(args...)
}

type logger struct {
	*log.Logger
}

func newLogger(l *log.Logger) *logger {
	return &logger{l}
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.Printf(format, args...)
}

func (l *logger) Debugln(args ...interface{}) {
	l.Println(args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.Printf(format, args...)
}

func (l *logger) Infoln(args ...interface{}) {
	l.Println(args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.Printf(format, args...)
}

func (l *logger) Warnln(args ...interface{}) {
	l.Println(args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.Printf(format, args...)
}

func (l *logger) Errorln(args ...interface{}) {
	l.Println(args...)
}
