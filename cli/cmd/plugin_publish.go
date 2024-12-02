package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/api"
	"github.com/cloudquery/cloudquery/cli/v6/internal/hub"
	"github.com/cloudquery/cloudquery/cli/v6/internal/publish"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	pluginPublishShort = "Publish to CloudQuery Hub."
	pluginPublishLong  = `Publish to CloudQuery Hub.

This publishes a plugin version to CloudQuery Hub from a local dist directory.
`
	pluginPublishExample = `
# Publish a plugin version from a local dist directory
cloudquery plugin publish`
)

func newCmdPluginPublish() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "publish [-D dist]",
		Short:   pluginPublishShort,
		Long:    pluginPublishLong,
		Example: pluginPublishExample,
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Set up a channel to listen for OS signals for graceful shutdown.
			ctx, cancel := context.WithCancel(cmd.Context())

			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGTERM)

			go func() {
				<-sigChan
				cancel()
			}()

			return runPluginPublish(ctx, cmd, args)
		},
	}
	cmd.Flags().StringP("dist-dir", "D", "dist", "Path to the dist directory")
	cmd.Flags().StringP("ui-dir", "U", "", "Path to the built plugin UI directory")
	cmd.Flags().BoolP("finalize", "f", false, `Finalize the plugin version after publishing. If false, the plugin version will be marked as draft.`)

	return cmd
}

func runPluginPublish(ctx context.Context, cmd *cobra.Command, args []string) error {
	tc := auth.NewTokenClient()
	token, err := tc.GetToken()
	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}

	uiDir := cmd.Flag("ui-dir").Value.String()
	if uiDir != "" {
		fi, err := os.Stat(uiDir)
		if err != nil {
			return fmt.Errorf("failed checking UI directory: %w", err)
		}
		if !fi.IsDir() {
			return errors.New("UI directory must be a directory")
		}
	}

	distDir := cmd.Flag("dist-dir").Value.String()
	pkgJSON, err := publish.ReadPackageJSON(distDir)
	if err != nil {
		return fmt.Errorf("failed to read package.json: %w", err)
	}

	if pkgJSON.Team != "" && pkgJSON.Name != "" && len(args) > 0 {
		log.Warn().Msgf("Passing team and plugin name as an argument is deprecated and no longer needed. Argument %q will be ignored", args[0])
	}

	teamName, pluginName := pkgJSON.Team, pkgJSON.Name
	if teamName == "" || pluginName == "" {
		parts := strings.Split(args[0], "/")
		if len(parts) != 2 {
			return errors.New("invalid plugin name. Must be in format <team_name>/<plugin_name>")
		}
		teamName, pluginName = parts[0], parts[1]
		pkgJSON.Team, pkgJSON.Name = teamName, pluginName
	}

	name := fmt.Sprintf("%s/%s@%s", teamName, pluginName, pkgJSON.Version)
	fmt.Printf("Publishing plugin %s to CloudQuery Hub...\n", name)

	c, err := api.NewClient(token.Value)
	if err != nil {
		return err
	}

	specJsonSchema, err := publish.GetSpecJsonScheme(distDir)
	if err != nil {
		return err
	}

	// create new draft version
	err = publish.CreateNewPluginDraftVersion(ctx, c, teamName, pluginName, pkgJSON, specJsonSchema)
	if err != nil {
		return fmt.Errorf("failed to create new draft version: %w", err)
	}

	if pkgJSON.Kind == cloudquery_api.PluginKindSource {
		// upload table schemas
		fmt.Println("Uploading table schemas...")
		tablesJSONPath := filepath.Join(distDir, "tables.json")
		err = publish.UploadTableSchemas(ctx, c, teamName, pluginName, tablesJSONPath, pkgJSON)
		if err != nil {
			return fmt.Errorf("failed to upload table schemas: %w", err)
		}
	}

	// upload docs
	fmt.Println("Uploading docs...")
	docsDir := filepath.Join(distDir, "docs")
	if err := publish.UploadPluginDocs(ctx, c, teamName, string(pkgJSON.Kind), pluginName, pkgJSON.Version, docsDir, true); err != nil {
		return fmt.Errorf("failed to upload docs: %w", err)
	}

	// upload binaries
	if err := publishPluginAssets(ctx, c, token.String(), distDir, pkgJSON); err != nil {
		return fmt.Errorf("failed to upload binaries: %w", err)
	}

	if uiDir != "" {
		fmt.Println("Uploading UI assets...")
		if err := publish.UploadPluginUIAssets(ctx, c, teamName, string(pkgJSON.Kind), pluginName, pkgJSON.Version, uiDir); err != nil {
			return fmt.Errorf("failed to upload UI assets: %w", err)
		}
	}

	// optional: mark plugin as draft=false
	finalize, err := cmd.Flags().GetBool("finalize")
	if err != nil {
		return err
	}

	if finalize {
		fmt.Println("Finalizing plugin version...")
		draft := false
		resp, err := c.UpdatePluginVersionWithResponse(ctx, teamName, pkgJSON.Kind, pluginName, pkgJSON.Version, cloudquery_api.UpdatePluginVersionJSONRequestBody{
			Draft: &draft,
		})
		if err != nil {
			return fmt.Errorf("failed to finalize plugin version: %w", err)
		}
		if resp.HTTPResponse.StatusCode > 299 {
			return hub.ErrorFromHTTPResponse(resp.HTTPResponse, resp)
		}
		fmt.Println("Success!")
		fmt.Printf("%s/%s@%s is now available on the CloudQuery Hub.\n", teamName, pluginName, pkgJSON.Version)
		return nil
	}

	fmt.Println("Success!")
	fmt.Println("\nNote: this plugin version is marked as draft=true. You can preview and finalize it on the CloudQuery Hub, or run `cloudquery plugin publish` with the --finalize flag.")

	return nil
}

func publishPluginAssets(ctx context.Context, c *cloudquery_api.ClientWithResponses, token, distDir string, pkgJSON publish.PackageJSONV1) error {
	if pkgJSON.PackageType == string(cloudquery_api.PluginPackageTypeDocker) {
		return publish.PublishToDockerRegistry(ctx, token, distDir, pkgJSON, publish.Opts{
			NoProgress: logConsole,
		})
	}

	return publish.PublishNativeBinaries(ctx, c, distDir, pkgJSON)
}
