package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery-api-go/config"
	"github.com/cloudquery/cloudquery/cli/v6/internal/env"
	"github.com/spf13/cobra"
)

const (
	platformOnboardShort   = "Set up your team's CloudQuery Platform destination."
	platformOnboardLong    = `Set up your team's CloudQuery Platform destination.

This provisions a platform tenant so you can sync data directly to the
CloudQuery Platform from your local CLI. Requires being enrolled in the
platform destination private beta.

The command will poll until the tenant is ready (usually 1-2 minutes).
`
	platformOnboardExample = `
# Onboard the currently selected team
cloudquery platform onboard

# Onboard a specific team
cloudquery platform onboard --team my-team
`
	defaultCloudAPIURL = "https://api.cloudquery.io"
	pollInterval       = 5 * time.Second
	pollTimeout        = 5 * time.Minute
)

type platformDestTenantResponse struct {
	PlatformURL string `json:"platform_url"`
	Status      string `json:"status"`
	TeamName    string `json:"team_name"`
}

func newCmdPlatformOnboard() *cobra.Command {
	var teamFlag string
	cmd := &cobra.Command{
		Use:     "onboard",
		Short:   platformOnboardShort,
		Long:    platformOnboardLong,
		Example: platformOnboardExample,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx, cancel := context.WithCancel(cmd.Context())
			defer cancel()

			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
			go func() {
				<-sigChan
				cancel()
			}()

			return runPlatformOnboard(ctx, teamFlag)
		},
	}
	cmd.Flags().StringVar(&teamFlag, "team", "", "Team name (defaults to currently selected team)")
	return cmd
}

func runPlatformOnboard(ctx context.Context, teamFlag string) error {
	tc := auth.NewTokenClient()
	tokenObj, err := tc.GetToken()
	if err != nil {
		return fmt.Errorf("not logged in. Run 'cloudquery login' first: %w", err)
	}
	token := tokenObj.Value

	teamName := teamFlag
	if teamName == "" {
		teamName, _ = config.GetValue("team")
	}
	if teamName == "" {
		return fmt.Errorf("no team selected. Use --team flag or run 'cloudquery switch' first")
	}

	apiURL := env.GetEnvOrDefault("CLOUDQUERY_API_URL", defaultCloudAPIURL)

	// 1. Call POST /teams/{team}/platform/onboard
	fmt.Fprintf(os.Stderr, "Setting up platform destination for team %q...\n", teamName)
	resp, err := platformAPICall(ctx, token, http.MethodPost, apiURL+"/teams/"+teamName+"/platform/onboard", nil)
	if err != nil {
		return err
	}

	if resp.Status == "active" {
		printOnboardSuccess(resp)
		return nil
	}

	// 2. Poll until active.
	fmt.Fprintf(os.Stderr, "Tenant provisioning in progress, polling for status...\n")
	deadline := time.After(pollTimeout)
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-deadline:
			fmt.Fprintf(os.Stderr, "\nProvisioning is taking longer than expected.\n")
			fmt.Fprintf(os.Stderr, "Run 'cloudquery platform onboard' again later to check status.\n")
			fmt.Fprintf(os.Stderr, "Platform URL (when ready): %s\n", resp.PlatformURL)
			return nil
		case <-ticker.C:
			status, err := platformAPICall(ctx, token, http.MethodGet, apiURL+"/teams/"+teamName+"/platform/status", nil)
			if err != nil {
				fmt.Fprintf(os.Stderr, ".")
				continue
			}
			if status.Status == "active" {
				fmt.Fprintln(os.Stderr)
				printOnboardSuccess(status)
				return nil
			}
			fmt.Fprintf(os.Stderr, ".")
		}
	}
}

func platformAPICall(ctx context.Context, token, method, url string, body any) (*platformDestTenantResponse, error) {
	var reqBody io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")
	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to reach CloudQuery API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden {
		return nil, fmt.Errorf("team is not enrolled in the platform destination beta. Contact support@cloudquery.io to request access")
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("authentication failed. Run 'cloudquery login' to refresh your token")
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("team has not been onboarded yet. Run 'cloudquery platform onboard' first")
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		snippet, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return nil, fmt.Errorf("CloudQuery API returned %d: %s", resp.StatusCode, strings.TrimSpace(string(snippet)))
	}

	var out platformDestTenantResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("failed to parse API response: %w", err)
	}
	return &out, nil
}

func printOnboardSuccess(resp *platformDestTenantResponse) {
	fmt.Fprintf(os.Stderr, "Platform destination ready for team %q!\n", resp.TeamName)
	fmt.Fprintf(os.Stderr, "Platform URL: %s\n\n", resp.PlatformURL)
	fmt.Fprintf(os.Stderr, "To sync data to the platform, add this to your CloudQuery config:\n\n")
	fmt.Fprintf(os.Stderr, "  kind: destination\n")
	fmt.Fprintf(os.Stderr, "  spec:\n")
	fmt.Fprintf(os.Stderr, "    name: platform\n")
	fmt.Fprintf(os.Stderr, "    path: cloudquery/platform\n")
	fmt.Fprintf(os.Stderr, "    spec:\n")
	fmt.Fprintf(os.Stderr, "      api_url: %s\n", resp.PlatformURL)
}
