package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery/cli/internal/auth"
	"github.com/spf13/cobra"
)

const (
	addonPublishShort = "Publish to CloudQuery Hub."
	addonPublishLong  = `Publish to CloudQuery Hub.

This publishes an addon version to CloudQuery Hub from a local dist directory.
`
	addonPublishExample = `
# Publish an addon version from a local dist directory
cloudquery addon publish my_team/my_addon`
)

func newCmdAddonPublish() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "publish <team_name>/<addon_name> [-D dist]",
		Short:   addonPublishShort,
		Long:    addonPublishLong,
		Example: addonPublishExample,
		Hidden:  true,
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

			return runAddonPublish(ctx, cmd, args)
		},
	}
	cmd.Flags().StringP("dist-dir", "D", "dist", "Path to the dist directory")
	cmd.Flags().BoolP("finalize", "f", false, `Finalize the plugin version after publishing. If false, the plugin version will be marked as draft=true.`)

	return cmd
}

type ManifestJSONV1 struct {
	Name    string `json:"name"`
	Version string `json:"version"`

	Message   string `json:"message"`
	PathToDoc string `json:"doc"`

	Type        string   `json:"type"`         // always "addon"
	AddonType   string   `json:"addon_type"`   // TODO unused
	AddonFormat string   `json:"addon_format"` // TODO unused
	PluginDeps  []string `json:"plugin_deps"`
	AddonDeps   []string `json:"addon_deps"`
}

func runAddonPublish(ctx context.Context, cmd *cobra.Command, args []string) error {
	token, err := auth.GetToken()
	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}

	distDir := cmd.Flag("dist-dir").Value.String()
	manifest, err := readManifestJSON(distDir)
	if err != nil {
		return fmt.Errorf("failed to read manifest.json: %w", err)
	}

	parts := strings.Split(args[0], "/")
	if len(parts) != 2 {
		return errors.New("invalid addon name. Must be in format <team_name>/<plugin_name>")
	}
	teamName, addonName := parts[0], parts[1]

	if manifest.Type != "addon" {
		return errors.New("manifest.json is not of addon type")
	}

	name := fmt.Sprintf("%s/%s@%s", teamName, addonName, manifest.Version)
	fmt.Printf("Publishing %s to CloudQuery Hub...\n", name)

	c, err := cloudquery_api.NewClientWithResponses(getEnvOrDefault("CLOUDQUERY_API_URL", defaultAPIURL),
		cloudquery_api.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			return nil
		}))
	if err != nil {
		return fmt.Errorf("failed to create hub client: %w", err)
	}

	// create new draft version
	err = createNewAddonDraftVersion(ctx, c, teamName, addonName, distDir, manifest)
	if err != nil {
		return fmt.Errorf("failed to create new draft version: %w", err)
	}

	// upload package
	fmt.Println("Uploading addon...")
	err = uploadAddon(ctx, c, teamName, addonName, distDir, manifest)
	if err != nil {
		return fmt.Errorf("failed to upload addon: %w", err)
	}

	// optional: mark plugin as draft=false
	finalize, err := cmd.Flags().GetBool("finalize")
	if err != nil {
		return err
	}

	if finalize {
		fmt.Println("Finalizing addon version...")
		draft := false
		resp, err := c.UpdateAddonVersionWithResponse(ctx, teamName, addonName, manifest.Version, cloudquery_api.UpdateAddonVersionJSONRequestBody{
			Draft: &draft,
		})
		if err != nil {
			return fmt.Errorf("failed to finalize addon version: %w", err)
		}
		if resp.HTTPResponse.StatusCode > 299 {
			return errorFromHTTPResponse(resp.HTTPResponse, resp)
		}
		fmt.Println("Success!")
		fmt.Printf("%s/%s@%s is now available on the CloudQuery Hub.\n", teamName, addonName, manifest.Version)
		return nil
	}

	fmt.Println("Success!")
	fmt.Println("\nNote: this addon version is marked as draft=true. You can preview and finalize it on the CloudQuery Hub, or run `cloudquery addon publish` with the --finalize flag.")

	return nil
}

func createNewAddonDraftVersion(ctx context.Context, c *cloudquery_api.ClientWithResponses, teamName, pluginName, distDir string, manifest ManifestJSONV1) error {
	doc := ""
	if manifest.PathToDoc != "" {
		b, err := os.ReadFile(filepath.Join(distDir, manifest.PathToDoc))
		if err != nil {
			return fmt.Errorf("failed to read doc file: %w", err)
		}
		doc = string(b)
	}

	body := cloudquery_api.CreateAddonVersionJSONRequestBody{
		Checksum: "", // TODO
		Doc:      doc,
		Message:  manifest.Message,

		AddonDeps:  &manifest.AddonDeps,
		PluginDeps: &manifest.PluginDeps,
	}
	resp, err := c.CreateAddonVersionWithResponse(ctx, teamName, pluginName, manifest.Version, body)
	if err != nil {
		return fmt.Errorf("failed to create addon version: %w", err)
	}
	if resp.HTTPResponse.StatusCode > 299 {
		err := errorFromHTTPResponse(resp.HTTPResponse, resp)
		if resp.HTTPResponse.StatusCode == http.StatusForbidden {
			return fmt.Errorf("%w. Hint: You may need to create the addon first", err)
		}
		return err
	}
	return nil
}

func uploadAddon(ctx context.Context, c *cloudquery_api.ClientWithResponses, teamName, addonName, localPath string, manifest ManifestJSONV1) error {
	resp, err := c.UploadAddonAssetWithResponse(ctx, teamName, addonName, manifest.Version)
	if err != nil {
		return fmt.Errorf("failed to upload addon: %w", err)
	}
	if resp.HTTPResponse.StatusCode > 299 {
		msg := fmt.Sprintf("failed to upload addon: %s", resp.HTTPResponse.Status)
		switch {
		case resp.JSON403 != nil:
			msg = fmt.Sprintf("%s: %s", msg, resp.JSON403.Message)
		case resp.JSON401 != nil:
			msg = fmt.Sprintf("%s: %s", msg, resp.JSON401.Message)
		}
		return fmt.Errorf(msg)
	}
	if resp.JSON201 == nil {
		return fmt.Errorf("upload response is nil, failed to upload addon")
	}
	uploadURL := resp.JSON201.Url

	if err := uploadFile(uploadURL, localPath); err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}
	return nil
}

func readManifestJSON(distDir string) (ManifestJSONV1, error) {
	v := SchemaVersion{}
	b, err := os.ReadFile(filepath.Join(distDir, "manifest.json"))
	if err != nil {
		return ManifestJSONV1{}, err
	}

	if err := json.Unmarshal(b, &v); err != nil {
		return ManifestJSONV1{}, err
	}
	if v.SchemaVersion != 1 {
		return ManifestJSONV1{}, errors.New("unsupported schema version. This CLI version only supports manifest.json v1. Try upgrading your CloudQuery CLI version")
	}

	manifest := ManifestJSONV1{}
	if err := json.Unmarshal(b, &manifest); err != nil {
		return ManifestJSONV1{}, err
	}
	return manifest, nil
}
