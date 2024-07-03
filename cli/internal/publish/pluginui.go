package publish

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery/cli/internal/hub"
)

func UploadPluginUIAssets(ctx context.Context, c *cloudquery_api.ClientWithResponses, teamName, pluginKind, pluginName, version, uiDir string) error {
	dirEntries, err := readFlatDir(uiDir)
	if err != nil {
		return fmt.Errorf("failed to read assets directory: %w", err)
	}

	assets := make([]cloudquery_api.PluginUIAssetUploadRequest, 0, len(dirEntries))
	urlPathVsDetails := make(map[string][]string, len(dirEntries))
	for _, dirEntry := range dirEntries {
		fullPath := filepath.Join(uiDir, dirEntry)
		urlPath := dirEntry
		if os.PathSeparator != '/' {
			urlPath = strings.ReplaceAll(urlPath, string(os.PathSeparator), "/")
		}

		filebytes := make([]byte, 512)
		{
			fp, err := os.Open(fullPath)
			if err != nil {
				return fmt.Errorf("failed to open file: %w", err)
			}
			if _, err := fp.Read(filebytes); err != nil {
				return fmt.Errorf("failed to read file: %w", err)
			}
			_ = fp.Close()
		}
		contentType := http.DetectContentType(filebytes)
		urlPathVsDetails[urlPath] = []string{fullPath, contentType}

		assets = append(assets, cloudquery_api.PluginUIAssetUploadRequest{
			Name:        urlPath,
			ContentType: &contentType,
		})
	}

	resp, err := c.UploadPluginUIAssetsWithResponse(
		ctx, teamName, cloudquery_api.PluginKind(pluginKind), pluginName, version,
		cloudquery_api.UploadPluginUIAssetsJSONRequestBody{
			Assets: assets,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to prepare for upload: %w", err)
	}
	if resp.HTTPResponse.StatusCode > 299 {
		return hub.ErrorFromHTTPResponse(resp.HTTPResponse, resp)
	}

	for _, asset := range resp.JSON201.Assets {
		details := urlPathVsDetails[asset.Name]
		if err := hub.UploadFileWithContentType(asset.UploadURL, details[0], details[1]); err != nil {
			return fmt.Errorf("failed to upload: %w", err)
		}
	}

	finalizeResp, err := c.FinalizePluginUIAssetUploadWithResponse(
		ctx, teamName, cloudquery_api.PluginKind(pluginKind), pluginName, version,
		cloudquery_api.FinalizePluginUIAssetUploadJSONRequestBody{
			UIID: resp.JSON201.UIID,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to finalize upload: %w", err)
	}
	if finalizeResp.HTTPResponse.StatusCode > 299 {
		return hub.ErrorFromHTTPResponse(resp.HTTPResponse, resp)
	}

	return nil
}

func readFlatDir(base string) (files []string, err error) {
	err = filepath.WalkDir(base, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
