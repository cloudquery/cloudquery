package logging

import (
	"io/ioutil"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestZerologKVAdapter_IsLevel(t *testing.T) {
	tests := []struct {
		name string
		lvl  zerolog.Level
	}{
		{
			name: "Trace",
			lvl:  zerolog.TraceLevel,
		},
		{
			name: "Debug",
			lvl:  zerolog.DebugLevel,
		},
		{
			name: "Info",
			lvl:  zerolog.InfoLevel,
		},
		{
			name: "Warn",
			lvl:  zerolog.WarnLevel,
		},
		{
			name: "Error",
			lvl:  zerolog.ErrorLevel,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := zerolog.New(ioutil.Discard)
			logger = logger.Level(tt.lvl)
			z := NewZHcLog(&logger, "")

			var isTrace, isDebug, isInfo, isWarn, isError bool
			if tt.lvl <= -1 {
				isTrace = true
			}
			if tt.lvl <= 0 {
				isDebug = true
			}
			if tt.lvl <= 1 {
				isInfo = true
			}
			if tt.lvl <= 2 {
				isWarn = true
			}
			if tt.lvl <= 3 {
				isError = true
			}
			assert.Equalf(t, isTrace, z.IsTrace(), "IsTrace()")
			assert.Equalf(t, isDebug, z.IsDebug(), "IsDebug()")
			assert.Equalf(t, isInfo, z.IsInfo(), "IsInfo()")
			assert.Equalf(t, isWarn, z.IsWarn(), "IsWarn()")
			assert.Equalf(t, isError, z.IsError(), "IsError()")
		})
	}
}
