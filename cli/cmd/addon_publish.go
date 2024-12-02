package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/api"
	"github.com/cloudquery/cloudquery/cli/v6/internal/hub"
	"github.com/cloudquery/cloudquery/cli/v6/internal/publish"
	"github.com/spf13/cobra"
)

const (
	addonPublishShort = "Publish to CloudQuery Hub."
	addonPublishLong  = `Publish to CloudQuery Hub.

This publishes an addon version to CloudQuery Hub from a manifest file.
`
	addonPublishExample = `
# Publish an addon version from a manifest file
cloudquery addon publish /path/to/manifest.json v1.0.0`
)

func newCmdAddonPublish() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "publish manifest.json v1.0.0 [--finalize]",
		Short:   addonPublishShort,
		Long:    addonPublishLong,
		Example: addonPublishExample,
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Set up a channel to listen for OS signals for graceful shutdown.
			ctx, cancel := context.WithCancel(cmd.Context())

			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGTERM)

			go func() {
				<-sigChan
				cancel()
			}()

			return runAddonPublish(ctx, cmd, args)
		},
	}
	cmd.Flags().BoolP("finalize", "f", false, `Finalize the addon version after publishing. If false, the addon version will be marked as draft.`)

	return cmd
}

func runAddonPublish(ctx context.Context, cmd *cobra.Command, args []string) error {
	tc := auth.NewTokenClient()
	token, err := tc.GetToken()
	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}

	manifestPath, version := args[0], args[1]
	manifestDir := filepath.Dir(manifestPath)

	manifest, err := publish.ReadManifestJSON(manifestPath)
	if err != nil {
		return fmt.Errorf("failed to read manifest.json: %w", err)
	}

	if manifest.Type != "addon" {
		return errors.New("manifest.json is not of addon type")
	}

	zipPath := filepath.Join(manifestDir, manifest.PathToZip)
	if _, err := os.Stat(zipPath); err != nil {
		return fmt.Errorf("could not read file %s: %w", zipPath, err)
	}

	name := fmt.Sprintf("%s/%s/%s@%s", manifest.TeamName, manifest.AddonType, manifest.AddonName, version)
	fmt.Printf("Publishing addon %s to CloudQuery Hub...\n", name)

	c, err := api.NewClient(token.Value)
	if err != nil {
		return err
	}

	// create new draft version
	if err := publish.CreateNewAddonDraftVersion(ctx, c, manifest, version, manifestDir, zipPath); err != nil {
		return fmt.Errorf("failed to create new draft version: %w", err)
	}

	// upload package
	fmt.Println("Uploading addon...")
	if err := publish.UploadAddon(ctx, c, manifest, version, zipPath); err != nil {
		return fmt.Errorf("failed to upload addon: %w", err)
	}

	// optional: mark addon as draft=false
	finalize, err := cmd.Flags().GetBool("finalize")
	if err != nil {
		return err
	}

	if finalize {
		fmt.Println("Finalizing addon version...")
		draft := false
		resp, err := c.UpdateAddonVersionWithResponse(ctx, manifest.TeamName, cloudquery_api.AddonType(manifest.AddonType), manifest.AddonName, version, cloudquery_api.UpdateAddonVersionJSONRequestBody{
			Draft: &draft,
		})
		if err != nil {
			return fmt.Errorf("failed to finalize addon version: %w", err)
		}
		if resp.HTTPResponse.StatusCode > 299 {
			return hub.ErrorFromHTTPResponse(resp.HTTPResponse, resp)
		}
		fmt.Println("Success!")
		fmt.Printf("%s/%s/%s@%s is now available on the CloudQuery Hub.\n", manifest.TeamName, manifest.AddonType, manifest.AddonName, version)
		return nil
	}

	fmt.Println("Success!")
	fmt.Println("\nNote: this addon version is marked as draft=true. You can preview and finalize it on the CloudQuery Hub, or run `cloudquery addon publish` with the --finalize flag.")

	return nil
}
