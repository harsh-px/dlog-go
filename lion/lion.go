package lion // import "go.pedge.io/dlog/lion"

import (
	"go.pedge.io/dlog"
	"go.pedge.io/lion"
)

func init() {
	dlog.SetLogger(lion.DefaultLogger)
}
