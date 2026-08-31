package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	neturl "net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	gosync "sync"
	"syscall"
	"time"

	"github.com/cenkalti/backoff/v7"
	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery-api-go/config"
	"github.com/cloudquery/cloudquery/cli/v6/internal/analytics"
	"github.com/cloudquery/cloudquery/cli/v6/internal/env"
	"github.com/cloudquery/cloudquery/cli/v6/internal/platform"
	"github.com/cloudquery/cloudquery/cli/v6/internal/team"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

const (
	// login command
	loginShort = "Login to CloudQuery Hub."
	loginLong  = `Login to CloudQuery Hub.

This is required to download plugins from CloudQuery Hub.

Local plugins and different registries don't need login.
`
	loginExample = `
# Log in to CloudQuery Hub
cloudquery login

# Log in to a specific team
cloudquery login --team my-team

# Log in directly to a CloudQuery Platform tenant
cloudquery login --host my-tenant.mycloudquery.com
`
)

func newCmdLogin() *cobra.Command {
	loginCmd := &cobra.Command{
		Use:     "login",
		Short:   loginShort,
		Long:    loginLong,
		Example: loginExample,
		Args:    cobra.MatchAll(cobra.ExactArgs(0), cobra.OnlyValidArgs),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Set up a channel to listen for OS signals for graceful shutdown.
			ctx, cancel := context.WithCancel(cmd.Context())

			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGTERM)

			go func() {
				<-sigChan
				cancel()
			}()

			return runLogin(ctx, cmd)
		},
	}
	loginCmd.Flags().StringP("team", "t", "", "Team to login to. Specify the team name, e.g. 'my-team' (not the display name)")
	loginCmd.Flags().String("host", "", "CloudQuery Platform tenant host to log in to directly, e.g. 'acme.mycloudquery.com', skipping the email-based routing")
	return loginCmd
}

// callbackHandler sends the browser to the success page of the app that
// authenticated it: cloud accounts, or the tenant origin passed with a platform
// token. Cloud's page is behind a cloud session a platform-only user lacks.
func callbackHandler(accountsURL string, onToken func(string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		onToken(token)

		if platform.IsPlatformToken(token) {
			if origin := tenantSuccessURL(r.URL.Query().Get("origin")); origin != "" {
				http.Redirect(w, r, origin, http.StatusSeeOther)
				return
			}
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			_, _ = io.WriteString(w, "You are signed in. You can close this window.\n")
			return
		}

		http.Redirect(w, r, accountsURL+"/success-close", http.StatusSeeOther)
	}
}

// tenantSuccessURL validates the origin, which is attacker-influenced: it
// arrives on a loopback URL anyone can open. Returns "" when untrusted.
func tenantSuccessURL(origin string) string {
	if origin == "" {
		return ""
	}
	u, err := neturl.Parse(origin)
	if err != nil || u.Host == "" || u.Path != "" || u.RawQuery != "" || u.Fragment != "" {
		return ""
	}
	switch {
	case u.Scheme == "https":
	case u.Scheme == "http" && isLoopbackHost(u.Hostname()):
	default:
		return ""
	}

	return u.String() + "/success-close"
}

func isLoopbackHost(host string) bool {
	if host == "localhost" {
		return true
	}
	ip := net.ParseIP(host)

	return ip != nil && ip.IsLoopback()
}

func waitForServer(ctx context.Context, url string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return err
	}

	_, retryErr := backoff.Retry(ctx, func() (any, error) {
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		_, _ = io.Copy(io.Discard, resp.Body)
		if resp.StatusCode == http.StatusOK {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to connect to local server. error code: %d", resp.StatusCode)
	}, backoff.WithBackOff(backoff.NewConstantBackOff(100*time.Millisecond)))
	return retryErr
}

func runLogin(ctx context.Context, cmd *cobra.Command) (err error) {
	accountsURL := env.GetEnvOrDefault("CLOUDQUERY_ACCOUNTS_URL", defaultAccountsURL)

	mux := http.NewServeMux()
	refreshToken := ""
	gotToken := make(chan struct{})
	var once gosync.Once
	mux.HandleFunc("/callback", callbackHandler(accountsURL, func(token string) {
		once.Do(func() {
			refreshToken = token
			close(gotToken)
		})
	}))
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	})
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	server := http.Server{
		Handler: mux,
		Addr:    listener.Addr().String(),
	}

	var serverErr error
	go func() {
		if err := server.Serve(listener); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				serverErr = fmt.Errorf("failed to serve: %w", err)
			}
		}
	}()
	localServerURL := "http://localhost:" + strconv.Itoa(listener.Addr().(*net.TCPAddr).Port)
	if err := waitForServer(ctx, localServerURL+"/health"); err != nil {
		return err
	}

	url := accountsURL + "?returnTo=" + localServerURL + "/callback"
	if host, _ := cmd.Flags().GetString("host"); host != "" {
		url = tenantLoginURL(host, localServerURL+"/callback")
	}
	if err := browser.OpenURL(url); err != nil {
		fmt.Printf("Failed to open browser. Please open %s manually and paste the token below:\n", url)

		stdinFd := int(os.Stdin.Fd())
		if !term.IsTerminal(stdinFd) {
			return errors.New("reading from non-terminal stdin is not supported. Hint: Consider setting an api key with the `CLOUDQUERY_API_KEY` env variable")
		}

		oldState, err := term.MakeRaw(stdinFd)
		if err != nil {
			return fmt.Errorf("failed setting stdin to raw mode: %w", err)
		}
		tty := term.NewTerminal(os.Stdin, "")
		refreshToken, err = tty.ReadLine()
		_ = term.Restore(stdinFd, oldState)

		if err != nil {
			return fmt.Errorf("failed to read token: %w", err)
		}

		refreshToken = strings.TrimSpace(refreshToken)
	} else {
		fmt.Printf("Opened browser at %s. Waiting for authentication to complete.\n", url)

		// Wait for an OS signal to begin shutting down.
		select {
		case <-ctx.Done():
			fmt.Println("Context cancelled. Shutting down server.")
		case <-gotToken:
		}
	}

	if refreshToken == "" {
		return errors.New("failed to get refresh token")
	}

	fmt.Println("Authenticating...")

	if platform.IsPlatformToken(refreshToken) {
		// A platform tenant completed the browser login: store the cqpd_ token
		// in its own file and take the team from its claims — this identity has
		// no Firebase refresh token to exchange and no Hub teams to list.
		if err := platform.SavePlatformToken(refreshToken); err != nil {
			return fmt.Errorf("failed to save platform token: %w", err)
		}
		_ = auth.RemoveRefreshToken()
		if err := setTeamOnPlatformLogin(cmd, refreshToken); err != nil {
			return err
		}
	} else {
		if err := platform.RemovePlatformToken(); err != nil {
			return fmt.Errorf("failed to remove previous platform token: %w", err)
		}

		err = auth.SaveRefreshToken(refreshToken)
		if err != nil {
			return fmt.Errorf("failed to save refresh token: %w", err)
		}

		tc := auth.NewTokenClient()
		token, err := tc.GetToken()
		if err != nil {
			return fmt.Errorf("failed to get auth token: %w", err)
		}

		if err = setTeamOnLogin(ctx, cmd, token.Value); err != nil {
			return fmt.Errorf("failed to set current team on login: %w", err)
		}
	}

	// Create a context for the shutdown with a 15-second timeout.
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	if err := server.Shutdown(ctxWithTimeout); err != nil {
		return fmt.Errorf("failed to shutdown server: %w", err)
	}

	if serverErr != nil {
		return serverErr
	}

	analytics.TrackLoginSuccess(ctx, invocationUUID.UUID)

	cmd.Println("✅ CLI successfully authenticated.")
	warnEnvCredentialsOverrideLogin(cmd)
	cmd.Println("Next, initialize your sync configuration:")
	cmd.Println(bold.Sprint("cloudquery init"))

	return nil
}

// tenantLoginURL points the browser straight at a platform tenant's login page,
// carrying the CLI callback the same way the accounts app forwards it.
func tenantLoginURL(host, callbackURL string) string {
	tenantURL := host
	if !strings.Contains(tenantURL, "://") {
		tenantURL = "https://" + tenantURL
	}
	return tenantURL + "/auth/login?cliReturnTo=" + neturl.QueryEscape(callbackURL)
}

// setTeamOnPlatformLogin sets the current team from the cqpd_ token's `tm`
// claim. Platform identities have no Hub teams to list, so the claim is the
// only source; a conflicting --team flag is an error rather than a silent
// mismatch.
func setTeamOnPlatformLogin(cmd *cobra.Command, token string) error {
	tokenTeam := platform.TeamFromToken(token)

	if cmd.Flags().Changed("team") {
		flagTeam := cmd.Flag("team").Value.String()
		if flagTeam != tokenTeam {
			return fmt.Errorf("the platform token belongs to team %q, not %q", tokenTeam, flagTeam)
		}
	}

	if tokenTeam == "" {
		return nil
	}

	if err := config.SetValue("team", tokenTeam); err != nil {
		return fmt.Errorf("failed to set team: %w", err)
	}
	if err := config.SetValue("team_internal", "false"); err != nil {
		return fmt.Errorf("failed to set team metadata: %w", err)
	}
	cmd.Printf("Your current team is set to %s.\n", tokenTeam)

	return nil
}

func setTeamOnLogin(ctx context.Context, cmd *cobra.Command, token string) error {
	cl, err := team.NewClient(token)
	if err != nil {
		return fmt.Errorf("failed to create API client: %w", err)
	}

	// list all available teams
	teams, err := cl.ListAllTeams(ctx)
	if err != nil {
		return fmt.Errorf("failed to list teams: %w", err)
	}

	if cmd.Flags().Changed("team") {
		// don't care about the current cached config value as the user explicitly passes the `team` flag
		currentTeam := cmd.Flag("team").Value.String()
		foundTeam, err := cl.FindTeam(teams, currentTeam)
		if err != nil {
			return fmt.Errorf("failed to validate team: %w", err)
		}
		err = config.SetValue("team", currentTeam)
		if err != nil {
			return fmt.Errorf("failed to set team: %w", err)
		}
		err = config.SetValue("team_internal", strconv.FormatBool(foundTeam.Internal))
		if err != nil {
			return fmt.Errorf("failed to set team metadata: %w", err)
		}
		return nil
	}

	currentTeam, err := config.GetValue("team")
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to get current team: %w", err)
	}
	foundTeam, err := cl.FindTeam(teams, currentTeam)
	if len(currentTeam) > 0 && err == nil {
		// The selected team is available to the user, so we're just using it
		cmd.Printf("Your current team is set to %s.\n", foundTeam.Name)

		// Make sure we update the internal flag in case it changed
		err = config.SetValue("team_internal", strconv.FormatBool(foundTeam.Internal))
		if err != nil {
			return fmt.Errorf("failed to set team metadata: %w", err)
		}
		return nil
	}

	// either the team is not set, or the currently selected team is unavailable
	switch len(teams) {
	case 0:
		// no available teams, urge the user to create one
		cmd.Printf("Your current team is not set.\n\n")
		cmd.Printf("There are no teams available to you.\n\n")
		cmd.Printf("Please create a team or accept an invite.")
		cmd.Printf("You should run `cloudquery switch <team>` to set your team afterwards.\n\n")
		if len(currentTeam) > 0 {
			// remove current team setting
			err = config.UnsetValue("team")
			if err != nil && !errors.Is(err, os.ErrNotExist) {
				return fmt.Errorf("failed to reset current team: %w", err)
			}
		}

	case 1:
		currentTeam = teams[0].Name
		err = config.SetValue("team", currentTeam)
		if err != nil {
			return fmt.Errorf("failed to set team: %w", err)
		}

		teamInternalStr := "false"
		if teams[0].Internal {
			teamInternalStr = "true"
		}

		err = config.SetValue("team_internal", teamInternalStr)
		if err != nil {
			return fmt.Errorf("failed to set team metadata: %w", err)
		}

		cmd.Printf("Your current team is set to %s.\n", currentTeam)

	default:
		cmd.Printf("Your current team is not set.\n\n")
		cmd.Printf("Teams available to you: %s\n\n", team.Names(teams))
		cmd.Printf("To set your current team, run `cloudquery switch <team>`\n\n")
		// we don't fail here immediately, as there are some additional commands the user can run in this state
	}

	return nil
}
