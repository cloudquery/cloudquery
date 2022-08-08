package logging

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_loggerWarningColor(t *testing.T) {
	ui.ColorWarning.EnableColor() // enable color forcefully for the test
	var b bytes.Buffer
	log := Configure(Config{
		ConsoleLoggingEnabled: true,
		Verbose:               true,
		console:               &b,
	})
	log.Warn().Msg("test")

	// skip first line
	out := b.String()
	pos := strings.Index(out, "\n")
	require.NotEqual(t, -1, pos, "can't find a linefeed in log output")
	out = out[pos+1:]

	// check the second field, log level
	fields := strings.Fields(out)
	require.GreaterOrEqual(t, len(fields), 2, "not enough fields in log output")
	got := fields[1]
	yellow := color.New(color.FgYellow)
	yellow.EnableColor()
	assert.Equal(t, yellow.Sprint("WRN"), got)
}
