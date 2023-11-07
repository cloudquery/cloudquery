package cmd

import (
	"context"
	"crypto/sha256"
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

	checksum := ""
	hc := *http.DefaultClient
	hc.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) == 1 && req.Response != nil {
			checksum = req.Response.Header.Get("X-Checksum-Sha256") // get checksum from first request
		}
		return nil
	}
	c, err := cloudquery_api.NewClientWithResponses(getEnvOrDefault(envAPIURL, defaultAPIURL),
		cloudquery_api.WithHTTPClient(&hc),
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

	res, err := c.DownloadAddonAsset(ctx, addonParts[0], cloudquery_api.AddonType(addonParts[1]), addonVer[0], addonVer[1])
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

	return downloadAddonFromResponse(res, checksum, target)
}

func downloadAddonFromResponse(res *http.Response, expectedChecksum, targetDir string) (retErr error) {
	var (
		fileWriter io.WriteCloser
		size       int64
		err        error
	)

	switch targetDir {
	case "-":
		fileWriter = os.Stdout
	default:
		zipPath := filepath.Join(targetDir, path.Base(res.Request.URL.Path))
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
	if expectedChecksum == "" {
		fmt.Fprintf(os.Stderr, "Warning - checksum not verified: %s\n", writtenChecksum)
	} else if writtenChecksum != expectedChecksum {
		return fmt.Errorf("checksum mismatch: expected %s, got %s", expectedChecksum, writtenChecksum)
	}

	return nil
}
