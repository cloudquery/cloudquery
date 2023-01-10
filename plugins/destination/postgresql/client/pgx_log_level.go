package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarn
	LogLevelInfo
	LogLevelDebug
	LogLevelTrace
)

func (r LogLevel) String() string {
	return [...]string{"error", "warn", "info", "debug", "trace"}[r]
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
	if *r, err = LogLevelFromString(loglevel); err != nil {
		return err
	}
	return nil
}

func LogLevelFromString(s string) (LogLevel, error) {
	switch s {
	case "trace":
		return LogLevelTrace, nil
	case "debug":
		return LogLevelDebug, nil
	case "info":
		return LogLevelInfo, nil
	case "warn":
		return LogLevelWarn, nil
	case "error":
		return LogLevelError, nil
	default:
		return LogLevelError, fmt.Errorf("invalid level %s", s)
	}
}
