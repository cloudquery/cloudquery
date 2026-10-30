package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemovedOptions(t *testing.T) {
	for _, tc := range []struct {
		name    string
		args    []string
		wantErr string
	}{
		{
			name:    "data-dir flag",
			args:    []string{"plugin", "install", "testdata/multiple-sources.yml", "--data-dir", "data"},
			wantErr: "unknown flag: --data-dir",
		},
		{
			name:    "color flag",
			args:    []string{"plugin", "install", "testdata/multiple-sources.yml", "--color", "on"},
			wantErr: "unknown flag: --color",
		},
		{
			name:    "top-level install command",
			args:    []string{"install", "testdata/multiple-sources.yml"},
			wantErr: "unknown command \"install\" for \"cloudquery\"\n\nDid you mean this?\n\tplugin\n",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			cmd := NewCmdRoot()
			cmd.SetArgs(append(tc.args, testCommandArgs(t)...))
			require.EqualError(t, cmd.Execute(), tc.wantErr)
		})
	}
}

func TestCQNoTelemetryIsRejected(t *testing.T) {
	t.Setenv("CQ_NO_TELEMETRY", "1")
	cmd := NewCmdRoot()
	cmd.SetArgs(append([]string{"plugin", "install", "testdata/multiple-sources.yml"}, testCommandArgs(t)...))
	require.EqualError(t, cmd.Execute(), "the CQ_NO_TELEMETRY environment variable is no longer supported, unset it and set CQ_TELEMETRY_LEVEL=none instead")
}
