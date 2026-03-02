package client

import (
	"context"
	"log/slog"

	"github.com/rs/zerolog"
)

// zerologSlogHandler bridges Go's slog to zerolog, mapping slog levels to zerolog levels.
type zerologSlogHandler struct {
	logger zerolog.Logger
}

func (h *zerologSlogHandler) Enabled(_ context.Context, level slog.Level) bool {
	var zlevel zerolog.Level
	switch {
	case level >= slog.LevelError:
		zlevel = zerolog.ErrorLevel
	case level >= slog.LevelWarn:
		zlevel = zerolog.WarnLevel
	case level >= slog.LevelInfo:
		zlevel = zerolog.InfoLevel
	default:
		zlevel = zerolog.DebugLevel
	}
	return h.logger.GetLevel() <= zlevel
}

func (h *zerologSlogHandler) Handle(_ context.Context, r slog.Record) error {
	var event *zerolog.Event
	switch {
	case r.Level >= slog.LevelError:
		event = h.logger.Error()
	case r.Level >= slog.LevelWarn:
		event = h.logger.Warn()
	case r.Level >= slog.LevelInfo:
		event = h.logger.Info()
	default:
		event = h.logger.Debug()
	}
	r.Attrs(func(a slog.Attr) bool {
		event = event.Any(a.Key, a.Value.Any())
		return true
	})
	event.Msg(r.Message)
	return nil
}

func (h *zerologSlogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	logger := h.logger
	for _, a := range attrs {
		logger = logger.With().Any(a.Key, a.Value.Any()).Logger()
	}
	return &zerologSlogHandler{logger: logger}
}

func (h *zerologSlogHandler) WithGroup(name string) slog.Handler {
	return h
}
