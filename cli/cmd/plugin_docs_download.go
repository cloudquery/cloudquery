package cmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/api"
	"github.com/cloudquery/cloudquery/cli/v6/internal/hub"
	"github.com/spf13/cobra"
)

const (
	pluginDocsDownloadShort = "Download plugin docs from CloudQuery Hub."
	pluginDocsDownloadLong  = `Download plugin docs from CloudQuery Hub.

This downloads documentation for a specific plugin version in CloudQuery hub to a local docs directory.
`
	pluginDocsDownloadExample = `
# Download plugin docs from CloudQuery Hub
cloudquery plugin docs download test-team/source/test-plugin@v1.0.0`
)

func newCmdPluginDocsDownload() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "download [-D docs] <team_name>/<plugin_kind>/<plugin_name>@<version>",
		Short:   pluginDocsDownloadShort,
		Long:    pluginDocsDownloadLong,
		Example: pluginDocsDownloadExample,
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

			return runPluginDocsDownload(ctx, cmd, args)
		},
	}
	cmd.Flags().StringP("docs-dir", "D", "docs", "Path to the docs directory")

	return cmd
}

func runPluginDocsDownload(ctx context.Context, cmd *cobra.Command, args []string) error {
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
		if !os.IsNotExist(err) {
			return err
		}
		if err := os.MkdirAll(docsDir, 0755); err != nil {
			return err
		}
	} else if !st.IsDir() {
		return fmt.Errorf("%s is not a directory", docsDir)
	}

	fmt.Printf("Downloading documentation for plugin %v...\n", pluginRef)

	c, err := api.NewClient(token.Value)
	if err != nil {
		return err
	}

	pg, ppg := cloudquery_api.Page(1), cloudquery_api.PerPage(1000)
	resp, err := c.ListPluginVersionDocsWithResponse(ctx, pluginRef.TeamName, cloudquery_api.PluginKind(pluginRef.Kind), pluginRef.Name, pluginRef.Version, &cloudquery_api.ListPluginVersionDocsParams{
		Page:    &pg,
		PerPage: &ppg,
	})
	if err != nil {
		return fmt.Errorf("failed to read docs: %w", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return hub.ErrorFromHTTPResponse(resp.HTTPResponse, resp)
	}
	if resp.JSON200 == nil {
		return errors.New("failed to read docs: nil response")
	}
	for _, item := range resp.JSON200.Items {
		safeName := strings.ReplaceAll(item.Name, string(filepath.Separator), "_") + ".md"
		fmt.Print("  ", safeName, " ")
		fn := filepath.Join(docsDir, safeName)
		fp, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
		if err != nil {
			return err
		}
		nb, err := fp.WriteString(item.Content)
		if err != nil {
			_ = os.Remove(fn)
			return err
		}
		if err := fp.Close(); err != nil {
			return err
		}
		fmt.Println(nb)
	}

	fmt.Println("Success!")

	return nil
}
