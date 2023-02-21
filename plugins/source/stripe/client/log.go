package client

import "github.com/rs/zerolog"

type LeveledLogger struct {
	zerolog.Logger
}

func (l *LeveledLogger) Debugf(format string, args ...any) {
	l.Logger.Debug().Msgf(format, args...)
}

func (l *LeveledLogger) Infof(format string, args ...any) {
	l.Logger.Info().Msgf(format, args...)
}

func (l *LeveledLogger) Warnf(format string, args ...any) {
	l.Logger.Warn().Msgf(format, args...)
}

func (l *LeveledLogger) Errorf(format string, args ...any) {
	l.Logger.Error().Msgf(format, args...)
}
