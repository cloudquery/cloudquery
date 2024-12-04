package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strings"
	"syscall"

	cqapiauth "github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/api"
	"github.com/cloudquery/cloudquery/cli/v6/internal/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/publish"
	"github.com/spf13/cobra"
)

const (
	addonDownloadShort = "Download addon from CloudQuery Hub."
	addonDownloadLong  = `Download addon from CloudQuery Hub.

This downloads an addon from CloudQuery Hub to local disk.
`
	addonDownloadExample = `
# Download an addon to local disk
cloudquery addon download <publisher>/<addon-type>/<addon-name>@v1.0.0

# Further example 
cloudquery addon download cloudquery/transformation/aws-compliance-premium@v1.9.0`
)

func newCmdAddonDownload() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "download addon-team/addon-type/addon-name@v1.0.0 [-t directory]",
		Short:   addonDownloadShort,
		Long:    addonDownloadLong,
		Example: addonDownloadExample,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Set up a channel to listen for OS signals for graceful shutdown.
			ctx, cancel := context.WithCancel(cmd.Context())

			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGTERM)

			go func() {
				<-sigChan
				cancel()
			}()

			return runAddonDownload(ctx, cmd, args)
		},
	}
	cmd.Flags().StringP("target", "t", ".", `Download to specified directory. Use - for stdout`)

	return cmd
}

func runAddonDownload(ctx context.Context, cmd *cobra.Command, args []string) error {
	tc := cqapiauth.NewTokenClient()
	token, err := tc.GetToken()
	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}

	addonParts := strings.Split(args[0], "/")
	if len(addonParts) != 3 {
		return fmt.Errorf("invalid addon ref: %s", args[0])
	}
	addonVer := strings.Split(addonParts[2], "@")
	if len(addonVer) != 2 {
		return fmt.Errorf("invalid addon ref %q: no version specified", args[0])
	}
	if !strings.HasPrefix(addonVer[1], "v") {
		return fmt.Errorf("invalid addon ref %q: version must start with 'v'", args[0])
	}

	c, err := api.NewClient(token.Value)
	if err != nil {
		return err
	}

	targetDir, err := cmd.Flags().GetString("target")
	if err != nil {
		return err
	}

	currentTeam, err := auth.GetTeamForToken(ctx, token)
	if err != nil {
		return err
	}

	location, checksum, err := publish.GetAddonMetadata(ctx, c, currentTeam, addonParts[0], addonParts[1], addonVer[0], addonVer[1])
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, location, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	// TODO: Remove this when https://github.com/timakin/bodyclose/issues/39 is fixed (false positive when close is in another package)
	// nolint: bodyclose
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make download request: %w", err)
	}
	if res.StatusCode > 399 {
		return fmt.Errorf("addon download failed: %d %s", res.StatusCode, location)
	}

	zipPath := "-"
	if targetDir != "-" {
		zipPath = filepath.Join(targetDir, path.Base(location))
	}

	return publish.DownloadAddonFromResponse(res, checksum, zipPath)
}
