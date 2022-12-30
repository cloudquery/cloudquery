package client

import "github.com/rs/zerolog"

type LeveledLogger struct {
	zerolog.Logger
}

func (l *LeveledLogger) Debugf(format string, args ...interface{}) {
	l.Logger.Debug().Msgf(format, args...)
}

func (l *LeveledLogger) Infof(format string, args ...interface{}) {
	l.Logger.Info().Msgf(format, args...)
}

func (l *LeveledLogger) Warnf(format string, args ...interface{}) {
	l.Logger.Warn().Msgf(format, args...)
}

func (l *LeveledLogger) Errorf(format string, args ...interface{}) {
	l.Logger.Error().Msgf(format, args...)
}
