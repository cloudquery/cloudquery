package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/api"
	"github.com/cloudquery/cloudquery/cli/v6/internal/hub"
	"github.com/cloudquery/cloudquery/cli/v6/internal/publish"
	"github.com/spf13/cobra"
)

const (
	pluginUIAssetsUploadShort = "Upload plugin UI assets to CloudQuery Hub."
	pluginUIAssetsUploadLong  = `Upload plugin UI assets to CloudQuery Hub.

This uploads UI assets for a specific plugin version from a local build directory to CloudQuery hub.
`
	pluginUIAssetsUploadExample = `
# Upload plugin UI assets to CloudQuery Hub
cloudquery plugin uiassets upload test-team/source/test-plugin@v1.0.0`
)

func newCmdPluginUIAssetsUpload() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "upload [-D build] <team_name>/<plugin_kind>/<plugin_name>@<version>",
		Short:   pluginUIAssetsUploadShort,
		Long:    pluginUIAssetsUploadLong,
		Example: pluginUIAssetsUploadExample,
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

			return runPluginUIAssetsUpload(ctx, cmd, args)
		},
	}
	cmd.Flags().StringP("assets-dir", "D", "build", "Path to the assets directory")

	return cmd
}

func runPluginUIAssetsUpload(ctx context.Context, cmd *cobra.Command, args []string) error {
	tc := auth.NewTokenClient()
	token, err := tc.GetToken()
	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}

	pluginRef, err := hub.ParseHubPluginRef(args[0])
	if err != nil {
		return err
	}

	uiAssetsDir := cmd.Flag("assets-dir").Value.String()
	st, err := os.Stat(uiAssetsDir)
	if err != nil {
		return err
	} else if !st.IsDir() {
		return fmt.Errorf("%s is not a directory", uiAssetsDir)
	}

	fmt.Printf("Uploading UI assets for plugin %v...\n", pluginRef)

	c, err := api.NewClient(token.Value)
	if err != nil {
		return err
	}

	if err := publish.UploadPluginUIAssets(ctx, c, pluginRef.TeamName, pluginRef.Kind, pluginRef.Name, pluginRef.Version, uiAssetsDir); err != nil {
		return fmt.Errorf("failed to upload UI assets: %w", err)
	}

	fmt.Println("Success!")
	return nil
}
