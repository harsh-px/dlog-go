[![CircleCI](https://circleci.com/gh/peter-edge/go-dlog/tree/master.png)](https://circleci.com/gh/peter-edge/go-dlog/tree/master)
[![API Documentation](http://img.shields.io/badge/api-Godoc-blue.svg?style=flat-square)](https://godoc.org/go.pedge.io/dlog)
[![GitHub License](https://img.shields.io/github/license/peter-edge/go-dlog.svg?style=flat-square)](https://github.com/peter-edge/go-dlog/blob/master/LICENSE)

dlog wraps common functionality for common golang logging packages.

The `dlog.Logger` interface wraps the common logging functionality. Every method on `dlog.Logger`
is also a global method on the `dlog` package. Given an implementation of `dlog.Logger`, you can
register it as the global logger by calling:

```go
func register(logger dlog.Logger) {
  dlog.SetLogger(logger)
}
```

To make things simple, packages for glog, logrus, and protolog are given with the ability to blank import:

```go
import (
  _ "go.pedge.io/dlog/glog" // set glog as the global logger
  _ "go.pedge.io/dlog/logrus" // set logrus as the global logger with default settings
  _ "go.pedge.io/dlog/protolog" // set protolog as the global logger with default settings
)
```

Or, do something more custom:

```go
func init() { // or anywhere
  logger := logrus.New()
  logger.Out = os.Stdout
  logger.Formatter = &logrus.TextFormatter{
    ForceColors: true,
  }
  dlog.SetLogger(logger)
}
```

By default, golang's standard logger is used.
