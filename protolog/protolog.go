package protolog // import "go.pedge.io/dlog/protolog"

import (
	"go.pedge.io/dlog"
	"go.pedge.io/protolog"
)

func init() {
	dlog.SetLogger(protolog.DefaultLogger)
}
