package cmd

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/spf13/cobra"
)

const (
	addonDownloadShort = "Download addon from CloudQuery Hub."
	addonDownloadLong  = `Download addon from CloudQuery Hub.

This downloads an addon from CloudQuery Hub to local disk.
`
	addonDownloadExample = `
# Download an addon to local disk
cloudquery addon download addon-team/addon-type/addon-name@v1.0.0`
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
	tc := auth.NewTokenClient()
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

	c, err := cloudquery_api.NewClientWithResponses(getEnvOrDefault(envAPIURL, defaultAPIURL),
		cloudquery_api.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			return nil
		}))
	if err != nil {
		return fmt.Errorf("failed to create hub client: %w", err)
	}

	target, err := cmd.Flags().GetString("target")
	if err != nil {
		return err
	}

	addon, addonVersion, err := getAddonMetadata(ctx, c, addonParts[0], addonParts[1], addonVer[0], addonVer[1])
	if err != nil {
		return fmt.Errorf("failed to get addon metadata: %w", err)
	}

	return downloadAddon(ctx, c, addon.TeamName, addon.AddonType, addon.Name, addonVersion.Name, addon.AddonFormat, addonVersion.Checksum, target)
}

func getAddonMetadata(ctx context.Context, c *cloudquery_api.ClientWithResponses, teamName, addonType, addonName, version string) (*cloudquery_api.Addon, *cloudquery_api.AddonVersion, error) {
	addonResp, err := c.GetAddonWithResponse(ctx, teamName, cloudquery_api.AddonType(addonType), addonName)
	if err != nil {
		if addonResp == nil {
			return nil, nil, fmt.Errorf("failed to get addon: %w", err)
		}
		return nil, nil, fmt.Errorf("failed to get addon: %w", errorFromHTTPResponse(addonResp.HTTPResponse, addonResp))
	}

	addonVersionResp, err := c.GetAddonVersionWithResponse(ctx, teamName, cloudquery_api.AddonType(addonType), addonName, version)
	if err != nil {
		if addonVersionResp == nil {
			return nil, nil, fmt.Errorf("failed to get addon version: %w", err)
		}
		return nil, nil, fmt.Errorf("failed to get addon version: %w", errorFromHTTPResponse(addonVersionResp.HTTPResponse, addonVersionResp))
	}
	return addonResp.JSON200, addonVersionResp.JSON200, nil
}

func addonFilename(teamName string, addonType cloudquery_api.AddonType, addonName, version string, format cloudquery_api.AddonFormat) string {
	return strings.Join([]string{teamName, string(addonType), addonName, version}, "_") + "." + string(format)
}

func downloadAddon(ctx context.Context, c *cloudquery_api.ClientWithResponses, teamName string, addonType cloudquery_api.AddonType, addonName, version string, format cloudquery_api.AddonFormat, checksum, targetDir string) (retErr error) {
	res, err := c.DownloadAddonAsset(ctx, teamName, addonType, addonName, version)
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}
	if res.StatusCode > 399 {
		resp, err := cloudquery_api.ParseDownloadAddonAssetResponse(res)
		if err != nil {
			return fmt.Errorf("failed to parse %d response: %w", res.StatusCode, err)
		}
		return fmt.Errorf("failed to download addon: %w", errorFromHTTPResponse(resp.HTTPResponse, resp))
	}

	var (
		fileWriter io.WriteCloser
		size       int64
	)

	switch targetDir {
	case "-":
		fileWriter = os.Stdout
	default:
		zipPath := filepath.Join(targetDir, addonFilename(teamName, addonType, addonName, version, format))
		if st, err := os.Stat(zipPath); err == nil {
			if st.IsDir() {
				return fmt.Errorf("file %s already exists: is a directory", zipPath)
			}
			return fmt.Errorf("file %s already exists", zipPath)
		}

		f, err := os.Create(zipPath)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		fileWriter = f

		defer func() {
			if retErr != nil {
				_ = os.Remove(zipPath)
				return
			}
			fmt.Fprintf(os.Stderr, "Wrote %d bytes to %s\n", size, zipPath)
		}()
	}

	shaWriter := sha256.New()
	w := io.MultiWriter(fileWriter, shaWriter)
	if size, err = io.Copy(w, res.Body); err != nil {
		_ = fileWriter.Close()
		return fmt.Errorf("failed to write: %w", err)
	}
	if err := fileWriter.Close(); err != nil {
		return fmt.Errorf("failed to close: %w", err)
	}
	if err := res.Body.Close(); err != nil {
		return fmt.Errorf("failed to close response body: %w", err)
	}

	writtenChecksum := fmt.Sprintf("%x", shaWriter.Sum(nil))
	if writtenChecksum != checksum {
		return fmt.Errorf("checksum mismatch: expected %s, got %s", checksum, writtenChecksum)
	}

	return nil
}
