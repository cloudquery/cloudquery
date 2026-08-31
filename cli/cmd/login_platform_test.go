package cmd

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	neturl "net/url"
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

func TestCallbackHandler(t *testing.T) {
	platformToken := mintPlatformTestToken(t, "acme")

	call := func(query string) *httptest.ResponseRecorder {
		var got string
		h := callbackHandler("https://accounts.cloudquery.io", func(token string) { got = token })
		w := httptest.NewRecorder()
		h(w, httptest.NewRequest(http.MethodGet, "/callback?"+query, nil))
		require.NotNil(t, got)

		return w
	}

	t.Run("platform token returns to the tenant success page", func(t *testing.T) {
		w := call("token=" + platformToken + "&origin=" + neturl.QueryEscape("https://acme.cloudquery.io"))

		require.Equal(t, http.StatusSeeOther, w.Code)
		require.Equal(t, "https://acme.cloudquery.io/success-close", w.Header().Get("Location"))
	})

	t.Run("loopback origin is allowed for local development", func(t *testing.T) {
		w := call("token=" + platformToken + "&origin=" + neturl.QueryEscape("http://localhost:4040"))

		require.Equal(t, http.StatusSeeOther, w.Code)
		require.Equal(t, "http://localhost:4040/success-close", w.Header().Get("Location"))
	})

	t.Run("untrusted origins fall back to plain text, never a redirect", func(t *testing.T) {
		for _, origin := range []string{
			"",
			"http://evil.example.com",
			"https://acme.cloudquery.io/path",
			"javascript:alert(1)",
			"//evil.example.com",
		} {
			w := call("token=" + platformToken + "&origin=" + neturl.QueryEscape(origin))

			require.Equal(t, http.StatusOK, w.Code, "origin %q", origin)
			require.Empty(t, w.Header().Get("Location"), "origin %q", origin)
			require.Contains(t, w.Body.String(), "You are signed in")
		}
	})

	t.Run("cloud token still redirects to accounts", func(t *testing.T) {
		w := call("token=firebase-refresh-token&origin=" + neturl.QueryEscape("https://acme.cloudquery.io"))

		require.Equal(t, http.StatusSeeOther, w.Code)
		require.Equal(t, "https://accounts.cloudquery.io/success-close", w.Header().Get("Location"))
	})
}
