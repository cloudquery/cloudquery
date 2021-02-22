package logging

import (
	"github.com/cloudquery/cloudquery/logging/keyvals"
	"github.com/hashicorp/go-hclog"
	"github.com/rs/zerolog"
	"io"
	"log"
	"reflect"
)

func NewZHcLog(l *zerolog.Logger, name string) hclog.Logger {
	return &ZerologKVAdapter{l, name, nil}
}

type ZerologKVAdapter struct {
	l    *zerolog.Logger
	name string

	impliedArgs []interface{}
}

func (z ZerologKVAdapter) Log(level hclog.Level, msg string, args ...interface{}) {
	switch level {
	case hclog.NoLevel:
		return
	case hclog.Trace:
		z.Trace(msg, args...)
	case hclog.Debug:
		z.Debug(msg, args...)
	case hclog.Info:
		z.Info(msg, args...)
	case hclog.Warn:
		z.Warn(msg, args...)
	case hclog.Error:
		z.Error(msg, args...)
	}
}

func (z ZerologKVAdapter) Trace(msg string, args ...interface{}) {
	z.l.Trace().Fields(keyvals.ToMap(args)).Msg(msg)
}

func (z ZerologKVAdapter) Debug(msg string, args ...interface{}) {
	z.l.Debug().Fields(keyvals.ToMap(args)).Msg(msg)
}

func (z ZerologKVAdapter) Info(msg string, args ...interface{}) {
	z.l.Info().Fields(keyvals.ToMap(args)).Msg(msg)
}

func (z ZerologKVAdapter) Warn(msg string, args ...interface{}) {
	z.l.Warn().Fields(keyvals.ToMap(args)).Msg(msg)
}

func (z ZerologKVAdapter) Error(msg string, args ...interface{}) {
	z.l.Error().Fields(keyvals.ToMap(args)).Msg(msg)
}

func (z ZerologKVAdapter) IsTrace() bool {
	return z.l.GetLevel() >= zerolog.TraceLevel
}

func (z ZerologKVAdapter) IsDebug() bool {
	return z.l.GetLevel() >= zerolog.DebugLevel
}

func (z ZerologKVAdapter) IsInfo() bool {
	return z.l.GetLevel() >= zerolog.InfoLevel
}

func (z ZerologKVAdapter) IsWarn() bool {
	return z.l.GetLevel() >= zerolog.WarnLevel
}

func (z ZerologKVAdapter) IsError() bool {
	return z.l.GetLevel() >= zerolog.ErrorLevel
}

func (z ZerologKVAdapter) ImpliedArgs() []interface{} {
	// Not supported
	return nil
}

func (z ZerologKVAdapter) With(args ...interface{}) hclog.Logger {

	l := z.l.With().Fields(keyvals.ToMap(args)).Logger()
	return NewZHcLog(&l, z.Name())
}

func (z ZerologKVAdapter) Name() string {
	return z.name
}

func (z ZerologKVAdapter) Named(name string) hclog.Logger {
	return NewZHcLog(z.l, name)
}

func (z ZerologKVAdapter) ResetNamed(name string) hclog.Logger {
	return &z
}

func (z *ZerologKVAdapter) SetLevel(level hclog.Level) {
	leveledLog := z.l.Level(convertLevel(level))
	z.l = &leveledLog
}

func (z ZerologKVAdapter) StandardLogger(opts *hclog.StandardLoggerOptions) *log.Logger {
	if opts == nil {
		opts = &hclog.StandardLoggerOptions{}
	}
	return log.New(z.StandardWriter(opts), "", 0)
}

func (z ZerologKVAdapter) StandardWriter(opts *hclog.StandardLoggerOptions) io.Writer {
	v := reflect.ValueOf(z.l)
	w := v.FieldByName("w")
	writer, ok := w.Interface().(zerolog.LevelWriter)
	if !ok {
		return nil
	}
	return writer
}

func convertLevel(level hclog.Level) zerolog.Level {
	switch level {
	case hclog.NoLevel:
		return zerolog.NoLevel
	case hclog.Trace:
		return zerolog.TraceLevel
	case hclog.Debug:
		return zerolog.DebugLevel
	case hclog.Info:
		return zerolog.InfoLevel
	case hclog.Warn:
		return zerolog.WarnLevel
	case hclog.Error:
		return zerolog.ErrorLevel
	}
	return zerolog.NoLevel
}
