package spec

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/invopop/jsonschema"
	"github.com/jackc/pgx/v5/tracelog"
)

type LogLevel tracelog.LogLevel

var logLevels = [...]string{
	tracelog.LogLevelTrace: tracelog.LogLevelTrace.String(),
	tracelog.LogLevelDebug: tracelog.LogLevelDebug.String(),
	tracelog.LogLevelInfo:  tracelog.LogLevelInfo.String(),
	tracelog.LogLevelWarn:  tracelog.LogLevelWarn.String(),
	tracelog.LogLevelError: tracelog.LogLevelError.String(),
	tracelog.LogLevelNone:  tracelog.LogLevelNone.String(),
}

func (LogLevel) JSONSchema() *jsonschema.Schema {
	levels := make([]any, len(logLevels))
	for i, lvl := range logLevels {
		levels[i] = lvl
	}

	return &jsonschema.Schema{
		Type:        "string",
		Enum:        levels,
		Title:       "PostgreSQL driver log level",
		Description: "Defines what [`pgx`](https://github.com/jackc/pgx) call events should be logged.",
		Default:     tracelog.LogLevelError.String(),
	}
}

func (r LogLevel) String() string {
	return r.LogLevel().String()
}

func (r LogLevel) LogLevel() tracelog.LogLevel {
	return tracelog.LogLevel(r)
}

func (r LogLevel) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(r.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (r *LogLevel) UnmarshalJSON(data []byte) (err error) {
	var loglevel string
	if err := json.Unmarshal(data, &loglevel); err != nil {
		return err
	}
	if *r, err = logLevelFromString(loglevel); err != nil {
		return err
	}
	return nil
}

func logLevelFromString(s string) (LogLevel, error) {
	if len(s) == 0 {
		return LogLevel(tracelog.LogLevelError), nil
	}
	res, err := tracelog.LogLevelFromString(strings.ToLower(s)) // we could use just the string value per enum in schema
	return LogLevel(res), err
}
