package client

import "github.com/rs/zerolog"

type saramaLoggerAdapter struct {
	logger zerolog.Logger
}

//nolint:revive
func NewSaramaLoggerAdapter(l zerolog.Logger) *saramaLoggerAdapter {
	return &saramaLoggerAdapter{
		logger: l,
	}
}

func (l *saramaLoggerAdapter) Print(v ...any) {
	l.logger.Info().Msgf("%v", v...)
}

func (l *saramaLoggerAdapter) Printf(format string, v ...any) {
	l.logger.Info().Msgf(format, v...)
}

func (l *saramaLoggerAdapter) Println(v ...any) {
	l.logger.Info().Msgf("%v", v)
}
