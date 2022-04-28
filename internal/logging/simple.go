package logging

import (
	"github.com/rs/zerolog"
)

type Logger interface {
	// Logf call this method to log regular messages about the
	// operations they perform.
	// Messages logged by this method are usually tagged with an `INFO` log
	// level in common logging libraries.
	Logf(format string, args ...interface{})

	// Errorf call this method to log errors they encounter while
	// sending events to the backend servers.
	// Messages logged by this method are usually tagged with an `ERROR` log
	// level in common logging libraries.
	Errorf(format string, args ...interface{})
}

// NewSimple Creates hclog.Logger adapter from a simpler logger interfaces
func NewSimple(l *zerolog.Logger, name string) Logger {
	return &ZerologSimpleLogger{l, name}
}

type ZerologSimpleLogger struct {
	l    *zerolog.Logger
	name string
}

func (z ZerologSimpleLogger) Logf(format string, args ...interface{}) {
	z.l.Info().Str("module", z.name).Msgf(format, args...)
}

func (z ZerologSimpleLogger) Errorf(format string, args ...interface{}) {
	z.l.Error().Str("module", z.name).Msgf(format, args...)
}
