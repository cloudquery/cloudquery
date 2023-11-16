package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
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
	cmd.Flags().Bool("sync", false, "Syncronize docs with CloudQuery Hub, deleting any docs that are not present locally")

	return cmd
}

func runPluginDocsUpload(ctx context.Context, cmd *cobra.Command, args []string) error {
	tc := auth.NewTokenClient()
	token, err := tc.GetToken()
	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}

	pluginRef, err := parseHubPluginRef(args[0])
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

	c, err := cloudquery_api.NewClientWithResponses(getEnvOrDefault(envAPIURL, defaultAPIURL),
		cloudquery_api.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			return nil
		}))
	if err != nil {
		return fmt.Errorf("failed to create hub client: %w", err)
	}

	remoteDocs := make(map[string]struct{})

	doSync, err := cmd.Flags().GetBool("sync")
	if err != nil {
		return err
	}

	if doSync {
		pg, ppg := cloudquery_api.Page(1), cloudquery_api.PerPage(1000)
		resp, err := c.ListPluginVersionDocsWithResponse(ctx, pluginRef.TeamName, cloudquery_api.PluginKind(pluginRef.Kind), pluginRef.Name, pluginRef.Version, &cloudquery_api.ListPluginVersionDocsParams{
			Page:    &pg,
			PerPage: &ppg,
		})
		if err != nil {
			return fmt.Errorf("failed to read docs: %w", err)
		}
		if resp.StatusCode() != http.StatusOK {
			return errorFromHTTPResponse(resp.HTTPResponse, resp)
		}
		if resp.JSON200 == nil {
			return fmt.Errorf("failed to read docs: nil response")
		}
		for _, item := range resp.JSON200.Items {
			remoteDocs[item.Name] = struct{}{}
		}
	}

	processed, err := uploadDocs(ctx, c, pluginRef.TeamName, pluginRef.Kind, pluginRef.Name, pluginRef.Version, docsDir)
	if err != nil {
		return fmt.Errorf("failed to upload docs: %w", err)
	}

	if doSync {
		for _, p := range processed {
			delete(remoteDocs, strings.TrimSuffix(p, ".md"))
		}
		if len(remoteDocs) > 0 {
			fmt.Println("Deleting docs...")
			if err := deleteDocs(ctx, c, pluginRef.TeamName, pluginRef.Kind, pluginRef.Name, pluginRef.Version, maps.Keys(remoteDocs)); err != nil {
				return fmt.Errorf("failed to delete docs: %w", err)
			}
		}
	}

	fmt.Println("Success!")
	return nil
}

func deleteDocs(ctx context.Context, c *cloudquery_api.ClientWithResponses, teamName, pluginKind, pluginName, version string, docs []string) error {
	resp, err := c.DeletePluginVersionDocsWithResponse(ctx, teamName, cloudquery_api.PluginKind(pluginKind), pluginName, version, cloudquery_api.DeletePluginVersionDocsJSONRequestBody{
		Names: docs,
	})
	if err != nil {
		return fmt.Errorf("failed to delete docs: %w", err)
	}
	if resp.HTTPResponse.StatusCode > 299 {
		return errorFromHTTPResponse(resp.HTTPResponse, resp)
	}
	return nil
}
