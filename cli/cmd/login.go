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
	"syscall"
	"time"

	"github.com/adrg/xdg"
	"github.com/cenkalti/backoff/v4"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

const (
	loginShort = "Login to CloudQuery Hub."
	loginLong  = `Login to CloudQuery Hub.

This is required to download plugins from CloudQuery Hub.

Local plugins and different registries don't need login.
`

	accountsURL = "https://accounts.cloudquery.io"
)

func newCmdLogin() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "login",
		Short:  loginShort,
		Long:   loginLong,
		Hidden: true,
		Args:   cobra.MatchAll(cobra.ExactArgs(0), cobra.OnlyValidArgs),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Set up a channel to listen for OS signals for graceful shutdown.
			ctx, cancel := context.WithCancel(cmd.Context())

			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGTERM)

			go func() {
				<-sigChan
				fmt.Println("Received SIGTERM!")
				cancel()
			}()

			return runLogin(ctx)
		},
	}
	return cmd
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

func runLogin(ctx context.Context) (err error) {
	tokenFilePath, err := xdg.CacheFile("cloudquery/token")
	if err != nil {
		return fmt.Errorf("can't determine a proper location for token file: %w", err)
	}
	tokenFile, err := os.OpenFile(tokenFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return fmt.Errorf("can't open token file %q for writing: %w", tokenFilePath, err)
	}
	defer func() {
		e := tokenFile.Close()
		if err == nil && e != nil {
			err = fmt.Errorf("can't close token file %q after writing: %w", tokenFilePath, e)
		}
	}()

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
	localServerAddr := fmt.Sprintf("http://localhost:%d", listener.Addr().(*net.TCPAddr).Port)
	if err := waitForServer(ctx, localServerAddr); err != nil {
		return err
	}

	url := fmt.Sprintf("%s?returnTo=%s/callback", accountsURL, localServerAddr)
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
	if _, err = tokenFile.WriteString(refreshToken); err != nil {
		return fmt.Errorf("failed to write token to %q: %w", tokenFilePath, err)
	}

	fmt.Println("CLI successfully authenticated.")

	return nil
}
