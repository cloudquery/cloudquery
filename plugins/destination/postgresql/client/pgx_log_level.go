package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarn
	LogLevelInfo
	LogLevelDebug
	LogLevelTrace
)

var logLevels = [...]string{
	LogLevelError: "error",
	LogLevelWarn:  "warn",
	LogLevelInfo:  "info",
	LogLevelDebug: "debug",
	LogLevelTrace: "trace",
}

func (r *LogLevel) String() string {
	return logLevels[*r]
}
func (r *LogLevel) MarshalJSON() ([]byte, error) {
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
		return LogLevelError, nil
	}

	s = strings.ToLower(s)

	for i, level := range logLevels {
		if level == s {
			return LogLevel(i), nil
		}
	}

	return LogLevelError, fmt.Errorf("invalid level %s", s)
}
