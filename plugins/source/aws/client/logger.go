package client

import (
	"github.com/rs/zerolog"
)

type awsLogger struct {
	logger zerolog.Logger
	accounts []string
}

func (a *awsLogger) Trace(msg string, args ...interface{}) {
	a.logger.Trace().Msgf(msg, args...)
}

func (a *awsLogger) Debug(msg string, args ...interface{}) {
	a.logger.Debug().Msgf(msg, args...)
}

func (a *awsLogger) Info(msg string, args ...interface{}) {
	a.logger.Debug().Msgf(msg, args...)
}

func (a *awsLogger) Warn(msg string, args ...interface{}) {
	a.logger.Warn().Msgf(msg, args...)
}

func (a *awsLogger) Error(msg string, args ...interface{}) {
	a.logger.Error().Msgf(msg, args...)
}
