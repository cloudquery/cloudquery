package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery-api-go/config"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/cloudquery/cloudquery/cli/internal/team"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
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
		Hidden:  true,
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
	apiURL := getEnvOrDefault("CLOUDQUERY_API_URL", defaultAPIURL)

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
		fmt.Printf("Failed to open browser at %s. Please open the URL manually.\n", accountsURL)
	} else {
		fmt.Printf("Opened browser at %s. Waiting for authentication to complete.\n", url)
	}

	// Wait for an OS signal to begin shutting down.
	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled. Shutting down server.")
	case <-gotToken:
	}

	// Create a context for the shutdown with a 5-second timeout.
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown server: %w", err)
	}

	if serverErr != nil {
		return serverErr
	}

	if refreshToken == "" {
		return fmt.Errorf("failed to get refresh token")
	}
	err = auth.SaveRefreshToken(refreshToken)
	if err != nil {
		return fmt.Errorf("failed to save refresh token: %w", err)
	}

	if cmd.Flags().Changed("team") {
		selectedTeam := cmd.Flag("team").Value.String()
		tc := auth.NewTokenClient()
		token, err := tc.GetToken()
		if err != nil {
			return fmt.Errorf("failed to get auth token: %w", err)
		}
		cl, err := team.NewClient(apiURL, token)
		if err != nil {
			return fmt.Errorf("failed to create API client: %w", err)
		}
		err = cl.ValidateTeam(ctx, selectedTeam)
		if err != nil {
			return fmt.Errorf("failed to set team: %w", err)
		}
		err = config.SetValue("team", selectedTeam)
		if err != nil {
			return fmt.Errorf("failed to set team: %w", err)
		}
	}

	cmd.Println("CLI successfully authenticated.")

	return nil
}
