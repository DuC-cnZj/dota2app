package adapter

import (
	"github.com/DuC-cnZj/dota2app/pkg/dlog"
	"strings"
)

type GinWriter struct{}

func (e *GinWriter) Write(p []byte) (n int, err error) {
	if strings.Index(string(p), "[GIN-debug]") == 0 {
		return 0, nil
	}

	dlog.Debug(string(p))

	return len(p), nil
}

