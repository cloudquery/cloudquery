package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery-api-go/config"
	"github.com/cloudquery/cloudquery/cli/internal/team"
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
	return loginCmd
}

func waitForServer(ctx context.Context, url string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return err
	}

	return backoff.Retry(func() error {
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		_, _ = io.Copy(io.Discard, resp.Body)
		if resp.StatusCode == http.StatusOK {
			return nil
		}
		return fmt.Errorf("failed to connect to local server. error code: %d", resp.StatusCode)
	}, backoff.WithContext(backoff.NewConstantBackOff(100*time.Millisecond), ctx))
}

func runLogin(ctx context.Context, cmd *cobra.Command) (err error) {
	accountsURL := getEnvOrDefault("CLOUDQUERY_ACCOUNTS_URL", defaultAccountsURL)
	apiURL := getEnvOrDefault(envAPIURL, defaultAPIURL)

	mux := http.NewServeMux()
	refreshToken := ""
	gotToken := make(chan struct{})
	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, accountsURL+"/success-close", http.StatusSeeOther)
		refreshToken = r.URL.Query().Get("token")
		close(gotToken)
	})
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
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
	if err := browser.OpenURL(url); err != nil {
		fmt.Printf("Failed to open browser. Please open %s manually and paste the token below:\n", accountsURL)

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

	err = auth.SaveRefreshToken(refreshToken)
	if err != nil {
		return fmt.Errorf("failed to save refresh token: %w", err)
	}

	tc := auth.NewTokenClient()
	token, err := tc.GetToken()
	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}
	cl, err := team.NewClient(apiURL, token.Value)
	if err != nil {
		return fmt.Errorf("failed to create API client: %w", err)
	}

	currentTeam, err := config.GetValue("team")
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to get current team: %w", err)
	}
	if cmd.Flags().Changed("team") {
		selectedTeam := cmd.Flag("team").Value.String()
		err = cl.ValidateTeam(ctx, selectedTeam)
		if err != nil {
			return fmt.Errorf("failed to set team: %w", err)
		}
		err = config.SetValue("team", selectedTeam)
		if err != nil {
			return fmt.Errorf("failed to set team: %w", err)
		}
	} else {
		if currentTeam == "" {
			// if current team is not set, try to set it from the API
			teams, err := cl.ListAllTeams(ctx)
			if err != nil {
				return fmt.Errorf("failed to list teams: %w", err)
			}
			if len(teams) == 1 {
				err = config.SetValue("team", teams[0])
				if err != nil {
					return fmt.Errorf("failed to set team: %w", err)
				}
				cmd.Printf("Your current team is set to %s.\n", teams[0])
			} else {
				cmd.Printf("Your current team is not set.\n\n")
				cmd.Printf("Teams available to you: " + strings.Join(teams, ", ") + "\n\n")
				cmd.Printf("To set your current team, run `cloudquery switch <team>`\n\n")
			}
		} else {
			cmd.Printf("Your current team is set to %s.\n", currentTeam)
		}
	}

	// Create a context for the shutdown with a 15-second timeout.
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown server: %w", err)
	}

	if serverErr != nil {
		return serverErr
	}

	cmd.Println("CLI successfully authenticated.")

	return nil
}
