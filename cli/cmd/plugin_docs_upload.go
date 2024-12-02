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
	pluginDocsUploadShort = "Upload plugin docs to CloudQuery Hub."
	pluginDocsUploadLong  = `Upload plugin docs to CloudQuery Hub.

This uploads documentation for a specific plugin version from a local docs directory to CloudQuery hub.
`
	pluginDocsUploadExample = `
# Upload plugin docs to CloudQuery Hub
cloudquery plugin docs upload test-team/source/test-plugin@v1.0.0`
)

func newCmdPluginDocsUpload() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "upload [-D docs] [--sync] <team_name>/<plugin_kind>/<plugin_name>@<version>",
		Short:   pluginDocsUploadShort,
		Long:    pluginDocsUploadLong,
		Example: pluginDocsUploadExample,
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

			return runPluginDocsUpload(ctx, cmd, args)
		},
	}
	cmd.Flags().StringP("docs-dir", "D", "docs", "Path to the docs directory")
	cmd.Flags().Bool("sync", false, "Synchronize docs with CloudQuery Hub, deleting any docs that are not present locally")

	return cmd
}

func runPluginDocsUpload(ctx context.Context, cmd *cobra.Command, args []string) error {
	tc := auth.NewTokenClient()
	token, err := tc.GetToken()
	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}

	pluginRef, err := hub.ParseHubPluginRef(args[0])
	if err != nil {
		return err
	}

	docsDir := cmd.Flag("docs-dir").Value.String()
	st, err := os.Stat(docsDir)
	if err != nil {
		return err
	} else if !st.IsDir() {
		return fmt.Errorf("%s is not a directory", docsDir)
	}

	fmt.Printf("Uploading documentation for plugin %v...\n", pluginRef)

	c, err := api.NewClient(token.Value)
	if err != nil {
		return err
	}

	doSync, err := cmd.Flags().GetBool("sync")
	if err != nil {
		return err
	}

	if err := publish.UploadPluginDocs(ctx, c, pluginRef.TeamName, pluginRef.Kind, pluginRef.Name, pluginRef.Version, docsDir, doSync); err != nil {
		return fmt.Errorf("failed to upload docs: %w", err)
	}

	fmt.Println("Success!")
	return nil
}
