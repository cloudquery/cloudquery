package cmd

import (
	"bytes"
	"testing"

	cqapiauth "github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/platform"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func captureWarnings(t *testing.T, warn func(*cobra.Command)) string {
	t.Helper()
	var buf bytes.Buffer
	cmd := &cobra.Command{}
	cmd.SetOut(&buf)
	warn(cmd)
	return buf.String()
}

func TestWarnEnvCredentialsOverrideLogin(t *testing.T) {
	t.Run("no env credentials -> silent", func(t *testing.T) {
		t.Setenv(cqapiauth.EnvVarCloudQueryAPIKey, "")
		t.Setenv(platform.EnvPlatformToken, "")
		require.Empty(t, captureWarnings(t, warnEnvCredentialsOverrideLogin))
	})
	t.Run("both set -> warns about each", func(t *testing.T) {
		t.Setenv(cqapiauth.EnvVarCloudQueryAPIKey, "cqt_x")
		t.Setenv(platform.EnvPlatformToken, "cqpd_x.y")
		out := captureWarnings(t, warnEnvCredentialsOverrideLogin)
		require.Contains(t, out, cqapiauth.EnvVarCloudQueryAPIKey)
		require.Contains(t, out, platform.EnvPlatformToken)
		require.Contains(t, out, "takes precedence over this login")
	})
}

func TestWarnEnvCredentialsSurviveLogout(t *testing.T) {
	t.Run("no env credentials -> silent", func(t *testing.T) {
		t.Setenv(cqapiauth.EnvVarCloudQueryAPIKey, "")
		t.Setenv(platform.EnvPlatformToken, "")
		require.Empty(t, captureWarnings(t, warnEnvCredentialsSurviveLogout))
	})
	t.Run("both set -> warns each still authenticates", func(t *testing.T) {
		t.Setenv(cqapiauth.EnvVarCloudQueryAPIKey, "cqt_x")
		t.Setenv(platform.EnvPlatformToken, "cqpd_x.y")
		out := captureWarnings(t, warnEnvCredentialsSurviveLogout)
		require.Contains(t, out, cqapiauth.EnvVarCloudQueryAPIKey+" is still set")
		require.Contains(t, out, platform.EnvPlatformToken+" is still set")
	})
}
