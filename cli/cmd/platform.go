package cmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	cqapi "github.com/cloudquery/cloudquery-api-go"
	cqapiauth "github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/env"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/spf13/cobra"
)

const (
	platformOnboardShort = "Onboard the current team to the CloudQuery Platform destination."
	platformOnboardLong  = `Onboard the current team to the CloudQuery Platform destination.

Provisions a platform tenant for the team (idempotent) and waits until the
tenant is active. Once onboarded, future ` + "`cloudquery sync`" + ` runs will
automatically route data to the platform tenant in addition to any
destinations defined in the spec.

Authenticates with the CLI's existing token (team API key or login token).
`
	platformOnboardExample = `
# Onboard the currently-active team
cloudquery platform onboard
`

	defaultCloudAPIURL  = "https://api.cloudquery.io"
	cloudAPIURLEnv      = "CLOUDQUERY_API_URL"
	onboardPollInterval = 2 * time.Second
	onboardPollTimeout  = 2 * time.Minute
)

func newCmdPlatform() *cobra.Command {
	platformCmd := &cobra.Command{
		Use:   "platform",
		Short: "CloudQuery Platform commands",
	}
	platformCmd.AddCommand(newCmdPlatformOnboard())
	return platformCmd
}

func newCmdPlatformOnboard() *cobra.Command {
	return &cobra.Command{
		Use:     "onboard",
		Short:   platformOnboardShort,
		Long:    platformOnboardLong,
		Example: platformOnboardExample,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runPlatformOnboard(cmd.Context())
		},
	}
}

func runPlatformOnboard(ctx context.Context) error {
	token, err := authToken()
	if err != nil {
		return err
	}
	teamName, err := auth.GetTeamForToken(ctx, token)
	if err != nil {
		return fmt.Errorf("get team for token: %w", err)
	}
	if teamName == "" {
		return errors.New("team is not set; run `cloudquery switch <team>` first")
	}

	base := env.GetEnvOrDefault(cloudAPIURLEnv, defaultCloudAPIURL)
	client, err := newCloudAPIClient(base, token.Value)
	if err != nil {
		return fmt.Errorf("build cloud API client: %w", err)
	}

	fmt.Printf("→ POST %s/teams/%s/platform/onboard\n", base, teamName)
	onboard, err := client.OnboardPlatformDestinationWithResponse(ctx, cqapi.TeamName(teamName))
	if err != nil {
		return fmt.Errorf("POST onboard: %w", err)
	}
	if onboard.HTTPResponse.StatusCode == http.StatusForbidden {
		return errors.New("team is not enrolled in the platform destination beta (403)")
	}
	if onboard.JSON200 == nil && onboard.JSON201 == nil {
		return fmt.Errorf("POST onboard returned %d", onboard.HTTPResponse.StatusCode)
	}
	tenant := onboard.JSON201
	if tenant == nil {
		tenant = onboard.JSON200
	}
	fmt.Printf("  tenant created (status=%s)\n", tenant.Status)

	fmt.Printf("→ polling %s/teams/%s/platform/status until status=active (timeout %s)\n",
		base, teamName, onboardPollTimeout)
	pollCtx, cancel := context.WithTimeout(ctx, onboardPollTimeout)
	defer cancel()
	for {
		status, err := client.GetPlatformDestinationStatusWithResponse(pollCtx, cqapi.TeamName(teamName))
		if err != nil {
			return fmt.Errorf("GET status: %w", err)
		}
		if status.JSON200 == nil {
			return fmt.Errorf("GET status returned %d", status.HTTPResponse.StatusCode)
		}
		fmt.Printf("  status=%s\n", status.JSON200.Status)
		if string(status.JSON200.Status) == "active" {
			fmt.Printf("\n✅ Platform tenant active.\n   URL: %s\n\n", status.JSON200.PlatformUrl)
			fmt.Printf("Future `cloudquery sync` runs will auto-route data to this tenant.\n")
			return nil
		}
		select {
		case <-pollCtx.Done():
			return fmt.Errorf("tenant did not reach active state within %s", onboardPollTimeout)
		case <-time.After(onboardPollInterval):
		}
	}
}

func authToken() (cqapiauth.Token, error) {
	tc := cqapiauth.NewTokenClient()
	tok, err := tc.GetToken()
	if err != nil {
		return cqapiauth.UndefinedToken, fmt.Errorf("get auth token: %w (hint: run `cloudquery login` or set CLOUDQUERY_API_KEY)", err)
	}
	return tok, nil
}

// newCloudAPIClient builds a cloudquery-api-go client wired through
// retryablehttp so 5xx / network blips don't bubble up to the user, and
// stamps Bearer auth on every request.
func newCloudAPIClient(base, token string) (*cqapi.ClientWithResponses, error) {
	rc := retryablehttp.NewClient()
	rc.Logger = nil
	rc.HTTPClient.Timeout = 10 * time.Second
	rc.RetryMax = 3
	rc.RetryWaitMin = 200 * time.Millisecond
	rc.RetryWaitMax = 2 * time.Second

	return cqapi.NewClientWithResponses(base,
		cqapi.WithHTTPClient(rc.StandardClient()),
		cqapi.WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("Accept", "application/json")
			return nil
		}),
	)
}
