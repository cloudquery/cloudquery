package cmd

import (
	"encoding/base64"
	"testing"

	"github.com/cloudquery/cloudquery-api-go/config"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func mintPlatformTestToken(t *testing.T, team string) string {
	t.Helper()
	payload := base64.RawURLEncoding.EncodeToString(
		[]byte(`{"u":"https://acme.mycloudquery.com","tm":"` + team + `"}`),
	)
	sig := base64.RawURLEncoding.EncodeToString([]byte("sig"))
	return "cqpd_" + payload + "." + sig
}

func newLoginTestCmd(t *testing.T) *cobra.Command {
	t.Helper()
	require.NoError(t, config.SetConfigHome(t.TempDir()))
	t.Cleanup(func() {
		require.NoError(t, config.UnsetConfigHome())
	})
	return newCmdLogin()
}

func TestTenantLoginURL(t *testing.T) {
	require.Equal(t,
		"https://acme.mycloudquery.com/auth/login?cliReturnTo=http%3A%2F%2Flocalhost%3A1234%2Fcallback",
		tenantLoginURL("acme.mycloudquery.com", "http://localhost:1234/callback"),
	)
	require.Equal(t,
		"http://localhost:3000/auth/login?cliReturnTo=http%3A%2F%2Flocalhost%3A1234%2Fcallback",
		tenantLoginURL("http://localhost:3000", "http://localhost:1234/callback"),
	)
}

func TestSetTeamOnPlatformLoginSetsTeamFromClaims(t *testing.T) {
	cmd := newLoginTestCmd(t)

	require.NoError(t, setTeamOnPlatformLogin(cmd, mintPlatformTestToken(t, "team-x")))

	team, err := config.GetValue("team")
	require.NoError(t, err)
	require.Equal(t, "team-x", team)

	teamInternal, err := config.GetValue("team_internal")
	require.NoError(t, err)
	require.Equal(t, "false", teamInternal)
}

func TestSetTeamOnPlatformLoginRejectsMismatchedTeamFlag(t *testing.T) {
	cmd := newLoginTestCmd(t)
	require.NoError(t, cmd.Flags().Set("team", "other-team"))

	err := setTeamOnPlatformLogin(cmd, mintPlatformTestToken(t, "team-x"))
	require.ErrorContains(t, err, `belongs to team "team-x"`)
}

func TestSetTeamOnPlatformLoginWithoutTeamClaimIsNoop(t *testing.T) {
	cmd := newLoginTestCmd(t)

	require.NoError(t, setTeamOnPlatformLogin(cmd, mintPlatformTestToken(t, "")))

	_, err := config.GetValue("team")
	require.Error(t, err)
}
