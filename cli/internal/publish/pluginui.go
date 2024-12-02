package publish

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery/cli/v6/internal/hub"
	"github.com/cloudquery/cloudquery/cli/v6/internal/publish/images"
	"golang.org/x/sync/errgroup"
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

		contentType, err := images.DetectContentType(fullPath)
		if err != nil {
			return err
		}
		urlPathVsDetails[urlPath] = []string{fullPath, contentType}

		assets = append(assets, cloudquery_api.PluginUIAssetUploadRequest{
			Name:        urlPath,
			ContentType: &contentType,
		})
	}

	if _, ok := urlPathVsDetails["index.html"]; !ok {
		return errors.New("index.html is required in the UI directory")
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

	if resp.JSON201 == nil {
		return errors.New("upload response is nil, failed")
	}

	eg, egCtx := errgroup.WithContext(ctx)
	eg.SetLimit(4)

	for _, asset := range resp.JSON201.Assets {
		asset := asset
		details := urlPathVsDetails[asset.Name]
		eg.Go(func() error {
			return hub.UploadFileWithContentType(egCtx, asset.UploadURL, details[0], details[1])
		})
	}

	if err := eg.Wait(); err != nil {
		return fmt.Errorf("failed to upload: %w", err)
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
	base = filepath.Dir(base + string(filepath.Separator))
	err = filepath.WalkDir(base, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, strings.TrimPrefix(path, base+string(os.PathSeparator)))
		}
		return nil
	})
	return files, err
}
