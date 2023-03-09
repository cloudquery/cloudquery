package client

import (
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/rs/zerolog"
)

type Logger struct {
	Base zerolog.Logger
}

var (
	_ gremlingo.Logger = (*Logger)(nil)
)

func (l Logger) Log(verbosity gremlingo.LogVerbosity, v ...any) {
	if len(v) == 1 {
		l.Logf(verbosity, "%v", v[0])
		return
	}

	l.Base.Trace().Any("data", v).Any("verbosity", verbosity).Msg("gremlingo log")
}

func (l Logger) Logf(verbosity gremlingo.LogVerbosity, format string, v ...any) {
	l.Base.Trace().Any("verbosity", verbosity).Msgf(format, v...)
}
