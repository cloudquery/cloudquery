package cmd

import (
	"encoding/base64"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	neturl "net/url"
	"os"
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

	tokenTeam, err := platformLoginTeam(cmd, mintPlatformTestToken(t, "team-x"))
	require.NoError(t, err)
	require.NoError(t, setTeamOnPlatformLogin(cmd, tokenTeam))

	team, err := config.GetValue("team")
	require.NoError(t, err)
	require.Equal(t, "team-x", team)

	teamInternal, err := config.GetValue("team_internal")
	require.NoError(t, err)
	require.Equal(t, "false", teamInternal)
}

func TestPlatformLoginTeamRejectsMismatchedTeamFlagBeforePersisting(t *testing.T) {
	cmd := newLoginTestCmd(t)
	require.NoError(t, cmd.Flags().Set("team", "other-team"))

	_, err := platformLoginTeam(cmd, mintPlatformTestToken(t, "team-x"))
	require.ErrorContains(t, err, `belongs to team "team-x"`)

	_, err = config.GetValue("team")
	require.Error(t, err)
}

func TestSetTeamOnPlatformLoginWithoutTeamClaimClearsStaleTeam(t *testing.T) {
	cmd := newLoginTestCmd(t)
	require.NoError(t, config.SetValue("team", "team-from-a-previous-cloud-login"))

	tokenTeam, err := platformLoginTeam(cmd, mintPlatformTestToken(t, ""))
	require.NoError(t, err)
	require.NoError(t, setTeamOnPlatformLogin(cmd, tokenTeam))

	team, err := config.GetValue("team")
	require.NoError(t, err)
	require.Empty(t, team)
}

func TestNewCallbackListenerBindsLoopbackOnly(t *testing.T) {
	listener, err := newCallbackListener()
	require.NoError(t, err)
	t.Cleanup(func() { _ = listener.Close() })

	addr := listener.Addr().(*net.TCPAddr)
	require.True(t, addr.IP.Equal(net.IPv4(127, 0, 0, 1)), "listening on %s", addr)
	require.NotZero(t, addr.Port)
}

func TestCloudLoginURLCarriesCallbackState(t *testing.T) {
	returnTo := callbackURL(&net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 1234}, "STATE")
	require.Equal(t, "http://127.0.0.1:1234/callback?state=STATE", returnTo)

	loginURL, err := neturl.Parse(cloudLoginURL("https://accounts.cloudquery.io", returnTo))
	require.NoError(t, err)
	require.Equal(t, "accounts.cloudquery.io", loginURL.Host)
	require.Equal(t, returnTo, loginURL.Query().Get("returnTo"))
}

func TestLoginCallbackRoundTrip(t *testing.T) {
	for _, tc := range []struct {
		name          string
		host          string
		loginHost     string
		returnToParam string
		token         string
		wantLocation  string
	}{
		{
			name:          "cloud login",
			loginHost:     "accounts.cloudquery.io",
			returnToParam: "returnTo",
			token:         "firebase-refresh-token",
			wantLocation:  "https://accounts.cloudquery.io/success-close",
		},
		{
			name:          "platform login with --host",
			host:          "acme.mycloudquery.com",
			loginHost:     "acme.mycloudquery.com",
			returnToParam: "cliReturnTo",
			token:         mintPlatformTestToken(t, "acme"),
			wantLocation:  "https://acme.mycloudquery.com/success-close",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			listener, err := newCallbackListener()
			require.NoError(t, err)
			var got string
			handler, loginURL := newLoginCallback("https://accounts.cloudquery.io", tc.host, listener.Addr(), func(token string) { got = token })
			srv := httptest.NewUnstartedServer(handler)
			srv.Listener = listener
			srv.Start()
			t.Cleanup(srv.Close)

			login, err := neturl.Parse(loginURL)
			require.NoError(t, err)
			require.Equal(t, tc.loginHost, login.Host)

			callback, err := neturl.Parse(login.Query().Get(tc.returnToParam))
			require.NoError(t, err)
			query := callback.Query()
			query.Set("token", tc.token)
			query.Set("origin", "https://"+tc.loginHost)
			callback.RawQuery = query.Encode()

			client := &http.Client{CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
			resp, err := client.Get(callback.String())
			require.NoError(t, err)
			t.Cleanup(func() { _ = resp.Body.Close() })

			require.Equal(t, http.StatusSeeOther, resp.StatusCode)
			require.Equal(t, tc.wantLocation, resp.Header.Get("Location"))
			require.Equal(t, tc.token, got)
		})
	}
}

func TestCallbackHandlerRejectsMissingOrWrongState(t *testing.T) {
	for _, query := range []string{
		"token=attacker-token",
		"token=attacker-token&state=",
		"token=attacker-token&state=WRONG",
		"token=attacker-token&state=STATE2",
	} {
		called := false
		h := callbackHandler("https://accounts.cloudquery.io", "", "STATE", func(string) { called = true })
		w := httptest.NewRecorder()
		stderr := captureStderr(t, func() {
			h(w, httptest.NewRequest(http.MethodGet, "/callback?"+query, nil))
		})

		require.Equal(t, http.StatusForbidden, w.Code, "query %q", query)
		require.Empty(t, w.Header().Get("Location"), "query %q", query)
		require.False(t, called, "query %q", query)
		require.Contains(t, stderr, "rejected a login callback with a missing or wrong state value", "query %q", query)
	}
}

func captureStderr(t *testing.T, fn func()) string {
	t.Helper()
	r, w, err := os.Pipe()
	require.NoError(t, err)
	oldStderr := os.Stderr
	os.Stderr = w
	defer func() { os.Stderr = oldStderr }()

	fn()
	require.NoError(t, w.Close())
	out, err := io.ReadAll(r)
	require.NoError(t, err)
	return string(out)
}

func TestCallbackHandler(t *testing.T) {
	platformToken := mintPlatformTestToken(t, "acme")

	call := func(token, origin string) *httptest.ResponseRecorder {
		var got string
		h := callbackHandler("https://accounts.cloudquery.io", "platform.example.com", "STATE", func(t string) { got = t })
		w := httptest.NewRecorder()
		h(w, httptest.NewRequest(http.MethodGet, "/callback?state=STATE&token="+token+"&origin="+neturl.QueryEscape(origin), nil))
		require.Equal(t, token, got)

		return w
	}

	t.Run("platform token returns to the tenant success page", func(t *testing.T) {
		w := call(platformToken, "https://acme.cloudquery.io")

		require.Equal(t, http.StatusSeeOther, w.Code)
		require.Equal(t, "https://acme.cloudquery.io/success-close", w.Header().Get("Location"))
	})

	t.Run("platform token returns to the tenant host given with --host", func(t *testing.T) {
		w := call(platformToken, "https://platform.example.com")

		require.Equal(t, http.StatusSeeOther, w.Code)
		require.Equal(t, "https://platform.example.com/success-close", w.Header().Get("Location"))
	})

	t.Run("loopback origin is allowed for local development", func(t *testing.T) {
		w := call(platformToken, "http://localhost:4040")

		require.Equal(t, http.StatusSeeOther, w.Code)
		require.Equal(t, "http://localhost:4040/success-close", w.Header().Get("Location"))
	})

	t.Run("user info in a trusted origin is dropped from the redirect", func(t *testing.T) {
		w := call(platformToken, "https://user:pass@acme.cloudquery.io")

		require.Equal(t, http.StatusSeeOther, w.Code)
		require.Equal(t, "https://acme.cloudquery.io/success-close", w.Header().Get("Location"))
	})

	t.Run("untrusted origins fall back to plain text, never a redirect", func(t *testing.T) {
		for _, origin := range []string{
			"",
			"http://evil.example.com",
			"https://evil.example.com",
			"https://cloudquery.io.evil.example.com",
			"http://acme.cloudquery.io",
			"https://acme.cloudquery.io/path",
			"javascript:alert(1)",
			"//evil.example.com",
			"https://acme.cloudquery.io@evil.example.com",
			"https://acme.cloudquery.io:443@evil.example.com",
			"https://evil.example.com\\.cloudquery.io",
			"https://evil.example.com%2F.cloudquery.io",
			"https://evil.example.com:443.cloudquery.io",
		} {
			w := call(platformToken, origin)

			require.Equal(t, http.StatusOK, w.Code, "origin %q", origin)
			require.Empty(t, w.Header().Get("Location"), "origin %q", origin)
			require.Contains(t, w.Body.String(), "You are signed in")
		}
	})

	t.Run("cloud token still redirects to accounts", func(t *testing.T) {
		w := call("firebase-refresh-token", "https://acme.cloudquery.io")

		require.Equal(t, http.StatusSeeOther, w.Code)
		require.Equal(t, "https://accounts.cloudquery.io/success-close", w.Header().Get("Location"))
	})
}
