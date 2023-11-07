package cmd

import (
	"archive/zip"
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strings"
	"syscall"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/spf13/cobra"
)

const (
	addonPublishShort = "Publish to CloudQuery Hub."
	addonPublishLong  = `Publish to CloudQuery Hub.

This publishes an addon version to CloudQuery Hub from a manifest file and directory.
`
	addonPublishExample = `
# Publish an addon version from a manifest file and directory
cloudquery addon publish /path/to/manifest.json /path/to/addon-dir v1.0.0`
)

func newCmdAddonPublish() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "publish manifest.json /path/to/directory v1.0.0 [--finalize]",
		Short:   addonPublishShort,
		Long:    addonPublishLong,
		Example: addonPublishExample,
		Args:    cobra.ExactArgs(3),
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
	cmd.Flags().BoolP("finalize", "f", false, `Finalize the addon version after publishing. If false, the addon version will be marked as draft=true.`)

	return cmd
}

type ManifestJSONV1 struct {
	Type        string `json:"type"` // always "addon"
	TeamName    string `json:"team_name"`
	AddonName   string `json:"addon_name"`
	AddonType   string `json:"addon_type"`
	AddonFormat string `json:"addon_format"` // unused

	PathToMessage string `json:"message"`
	PathToDoc     string `json:"doc"`

	PluginDeps []string `json:"plugin_deps"`
	AddonDeps  []string `json:"addon_deps"`
}

func runAddonPublish(ctx context.Context, cmd *cobra.Command, args []string) error {
	tc := auth.NewTokenClient()
	token, err := tc.GetToken()
	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}

	manifestPath, addonPath, version := args[0], args[1], args[2]
	manifestDir := filepath.Dir(manifestPath)

	manifest, err := readManifestJSON(manifestPath)
	if err != nil {
		return fmt.Errorf("failed to read manifest.json: %w", err)
	}

	if manifest.Type != "addon" {
		return errors.New("manifest.json is not of addon type")
	}
	if manifest.AddonFormat != string(cloudquery_api.Zip) {
		return fmt.Errorf("unsupported manifest.json addon_format: only `%v` is supported", cloudquery_api.Zip)
	}

	if st, err := os.Stat(addonPath); err != nil {
		return fmt.Errorf("could not read directory %s: %w", addonPath, err)
	} else if !st.IsDir() {
		return fmt.Errorf("%s is not a directory", addonPath)
	}

	name := fmt.Sprintf("%s/%s/%s@%s", manifest.TeamName, manifest.AddonType, manifest.AddonName, version)
	fmt.Printf("Publishing addon %s to CloudQuery Hub...\n", name)

	c, err := cloudquery_api.NewClientWithResponses(getEnvOrDefault(envAPIURL, defaultAPIURL),
		cloudquery_api.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			return nil
		}))
	if err != nil {
		return fmt.Errorf("failed to create hub client: %w", err)
	}

	wrapDirectory := strings.Join([]string{manifest.TeamName, manifest.AddonType, manifest.AddonName, version}, "_")
	tempFile, checksum, err := zipAddon(addonPath, wrapDirectory, name)
	if err != nil {
		return fmt.Errorf("failed to zip: %w", err)
	}
	defer os.Remove(tempFile)

	// create new draft version
	if err := createNewAddonDraftVersion(ctx, c, manifest, version, manifestDir, checksum); err != nil {
		return fmt.Errorf("failed to create new draft version: %w", err)
	}

	// upload package
	fmt.Println("Uploading addon...")
	if err := uploadAddon(ctx, c, manifest, version, tempFile); err != nil {
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
			return errorFromHTTPResponse(resp.HTTPResponse, resp)
		}
		fmt.Println("Success!")
		fmt.Printf("%s/%s/%s@%s is now available on the CloudQuery Hub.\n", manifest.TeamName, manifest.AddonType, manifest.AddonName, version)
		return nil
	}

	fmt.Println("Success!")
	fmt.Println("\nNote: this addon version is marked as draft=true. You can preview and finalize it on the CloudQuery Hub, or run `cloudquery addon publish` with the --finalize flag.")

	return nil
}

func createNewAddonDraftVersion(ctx context.Context, c *cloudquery_api.ClientWithResponses, manifest ManifestJSONV1, version, manifestDir, checksum string) error {
	if manifest.PluginDeps == nil {
		manifest.PluginDeps = []string{}
	}
	if manifest.AddonDeps == nil {
		manifest.AddonDeps = []string{}
	}
	body := cloudquery_api.CreateAddonVersionJSONRequestBody{
		AddonDeps:  &manifest.AddonDeps,
		PluginDeps: &manifest.PluginDeps,
		Checksum:   checksum,
	}

	if manifest.PathToDoc != "" {
		b, err := os.ReadFile(filepath.Join(manifestDir, manifest.PathToDoc))
		if err != nil {
			return fmt.Errorf("failed to read doc file: %w", err)
		}
		body.Doc = string(b)
	}

	if manifest.PathToMessage != "" {
		b, err := os.ReadFile(filepath.Join(manifestDir, manifest.PathToMessage))
		if err != nil {
			return fmt.Errorf("failed to read message file: %w", err)
		}
		body.Message = string(b)
	}

	resp, err := c.CreateAddonVersionWithResponse(ctx, manifest.TeamName, cloudquery_api.AddonType(manifest.AddonType), manifest.AddonName, version, body)
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

func uploadAddon(ctx context.Context, c *cloudquery_api.ClientWithResponses, manifest ManifestJSONV1, version, localPath string) error {
	resp, err := c.UploadAddonAssetWithResponse(ctx, manifest.TeamName, cloudquery_api.AddonType(manifest.AddonType), manifest.AddonName, version)
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

func readManifestJSON(manifestPath string) (ManifestJSONV1, error) {
	v := SchemaVersion{}
	b, err := os.ReadFile(manifestPath)
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

func zipAddon(addonPath, wrapDirectory, comment string) (string, string, error) {
	s := sha256.New()
	zipFile, err := os.CreateTemp("", "cq-addon*.zip")
	if err != nil {
		return "", "", fmt.Errorf("failed to create temp file: %w", err)
	}

	w := io.MultiWriter(zipFile, s)
	zr := zip.NewWriter(w)
	if err := filepath.WalkDir(addonPath, func(pth string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Type() == os.ModeDir {
			return nil
		}
		fn := path.Join(wrapDirectory, filepath.ToSlash(strings.TrimPrefix(pth, addonPath)))
		zf, err := zr.Create(fn)
		if err != nil {
			return err
		}
		rd, err := os.Open(pth)
		if err != nil {
			return err
		}
		if _, err := io.Copy(zf, rd); err != nil {
			return err
		}
		if err := rd.Close(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return "", "", fmt.Errorf("failed to zip: %w", err)
	}

	if err := zr.SetComment(comment); err != nil {
		return "", "", fmt.Errorf("failed to set comment: %w", err)
	}
	if err := zr.Close(); err != nil {
		return "", "", fmt.Errorf("failed to close zip writer: %w", err)
	}
	return zipFile.Name(), fmt.Sprintf("%x", s.Sum(nil)), nil
}
