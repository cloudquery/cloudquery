package publish

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery/cli/v6/internal/hub"
	"github.com/cloudquery/cloudquery/cli/v6/internal/publish/images"
	"github.com/samber/lo"
	"golang.org/x/sync/errgroup"
)

const uiAssetBundleTarName = "bundle.tar.gz"

func UploadPluginUIAssets(ctx context.Context, c *cloudquery_api.ClientWithResponses, teamName, pluginKind, pluginName, version, uiDir string) error {
	dirEntries, err := readFlatDir(uiDir)
	if err != nil {
		return fmt.Errorf("failed to read assets directory: %w", err)
	}

	assets := make([]cloudquery_api.PluginUIAssetUploadRequest, 0, len(dirEntries))
	urlPathVsDetails := make(map[string][2]string, len(dirEntries))
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
		urlPathVsDetails[urlPath] = [2]string{fullPath, contentType}

		assets = append(assets, cloudquery_api.PluginUIAssetUploadRequest{
			Name:        urlPath,
			ContentType: &contentType,
		})
	}

	if _, ok := urlPathVsDetails[uiAssetBundleTarName]; ok {
		return fmt.Errorf("%s is a reserved name and cannot be used as an asset name", uiAssetBundleTarName)
	}

	if _, ok := urlPathVsDetails["index.html"]; !ok {
		return errors.New("index.html is required in the UI directory")
	}

	assets = append(assets, cloudquery_api.PluginUIAssetUploadRequest{
		Name:        uiAssetBundleTarName,
		ContentType: lo.ToPtr("application/gzip"),
	})

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

	var bundleAsset cloudquery_api.PluginUIAsset
	for _, asset := range resp.JSON201.Assets {
		if asset.Name == uiAssetBundleTarName {
			bundleAsset = asset
			continue
		}

		asset := asset
		details := urlPathVsDetails[asset.Name]
		eg.Go(func() error {
			return hub.UploadFileWithContentType(egCtx, asset.UploadURL, details[0], details[1])
		})
	}

	if bundleAsset.UploadURL == "" {
		return errors.New("bundle asset URL not found in the response")
	}

	bundleReader, bundleWriter := io.Pipe()
	defer bundleWriter.Close()

	eg.Go(func() error {
		defer bundleWriter.Close()
		return bundleTarGz(uiDir, bundleWriter)
	})

	eg.Go(func() error {
		return hub.UploadReaderWithContentType(egCtx, bundleAsset.UploadURL, bundleReader, "application/gzip")
	})

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

func bundleTarGz(uiDir string, bundleFile io.Writer) error {
	gw := gzip.NewWriter(bundleFile)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	fsys := os.DirFS(uiDir)

	return fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		file, err := fsys.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		info, err := file.Stat()
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return fmt.Errorf("failed to create tar file info header: %w", err)
		}

		header.Name = path

		err = tw.WriteHeader(header)
		if err != nil {
			return fmt.Errorf("failed to write tar header: %w", err)
		}

		if d.IsDir() {
			return nil
		}

		if _, err := io.Copy(tw, file); err != nil {
			return fmt.Errorf("failed to write file into tar: %w", err)
		}

		return nil
	})
}
