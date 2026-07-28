package cmd

import (
	"os"

	cqapiauth "github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/platform"
	"github.com/spf13/cobra"
)

// Env credentials silently outrank the on-disk login: GetToken prefers
// CLOUDQUERY_API_KEY over the saved refresh token, and the platform paths
// prefer CQ_PLATFORM_TOKEN. login/logout call these so the user learns when
// the login state they just changed is not what will actually authenticate.

func warnEnvCredentialsOverrideLogin(cmd *cobra.Command) {
	if os.Getenv(cqapiauth.EnvVarCloudQueryAPIKey) != "" {
		cmd.Printf("Warning: %s is set in your environment and takes precedence over this login.\n", cqapiauth.EnvVarCloudQueryAPIKey)
	}
	if os.Getenv(platform.EnvPlatformToken) != "" {
		cmd.Printf("Warning: %s is set in your environment; platform syncs and plugin downloads will use it instead of this login.\n", platform.EnvPlatformToken)
	}
}

func warnEnvCredentialsSurviveLogout(cmd *cobra.Command) {
	if os.Getenv(cqapiauth.EnvVarCloudQueryAPIKey) != "" {
		cmd.Printf("Warning: %s is still set in your environment and will continue to authenticate CLI commands.\n", cqapiauth.EnvVarCloudQueryAPIKey)
	}
	if os.Getenv(platform.EnvPlatformToken) != "" {
		cmd.Printf("Warning: %s is still set in your environment and will continue to authenticate platform syncs and plugin downloads.\n", platform.EnvPlatformToken)
	}
}
