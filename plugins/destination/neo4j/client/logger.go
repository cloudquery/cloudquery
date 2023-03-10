package client

import (
	neo4jlog "github.com/neo4j/neo4j-go-driver/v5/neo4j/log"
	"github.com/rs/zerolog"
)

type Logger struct {
	Base zerolog.Logger
}

var (
	_ neo4jlog.Logger     = (*Logger)(nil)
	_ neo4jlog.BoltLogger = (*Logger)(nil)
)

func (l Logger) Error(name string, id string, err error) {
	l.Base.Error().Str("neo4j_name", name).Str("neo4j_id", id).Err(err).Msg("neo4j error")
}

func (l Logger) Warnf(name string, id string, msg string, args ...any) {
	l.Base.Warn().Str("neo4j_name", name).Str("neo4j_id", id).Msgf(msg, args...)
}

func (l Logger) Infof(name string, id string, msg string, args ...any) {
	l.Base.Info().Str("neo4j_name", name).Str("neo4j_id", id).Msgf(msg, args...)
}

func (l Logger) Debugf(name string, id string, msg string, args ...any) {
	l.Base.Debug().Str("neo4j_name", name).Str("neo4j_id", id).Msgf(msg, args...)
}

func (l Logger) LogClientMessage(id string, msg string, args ...any) {
	l.Base.Trace().Str("neo4j_bolt_log", "client").Str("neo4j_bolt_id", id).Msgf(msg, args...)
}

func (l Logger) LogServerMessage(id string, msg string, args ...any) {
	l.Base.Trace().Str("neo4j_bolt_log", "server").Str("neo4j_bolt_id", id).Msgf(msg, args...)
}
