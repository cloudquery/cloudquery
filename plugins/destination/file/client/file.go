package client

import (
	"context"
	"io"
	"path"

	"github.com/cloudquery/cloudquery/plugins/destination/file/internal/backends/gcs"
	"github.com/cloudquery/cloudquery/plugins/destination/file/internal/backends/local"
	"github.com/cloudquery/cloudquery/plugins/destination/file/internal/backends/s3"
)

func (c *Client) openReadOnly(ctx context.Context, name string) (io.ReadCloser, error) {
	switch c.csvSpec.Backend {
	case BackendTypeLocal:
		name := path.Join(c.csvSpec.Directory, name)
		return local.OpenReadOnly(name)
	case BackendTypeGCS:
		name := path.Join(c.path, name)
		return gcs.OpenReadOnly(ctx, c.gcpStorageClient, c.baseDir, name)
	case BackendTypeS3:
		name := path.Join(c.path, name)
		return s3.OpenReadOnly(ctx, c.awsDownloader, c.baseDir, name)
	default:
		panic("unknown backend " + c.csvSpec.Backend.String())
	}
}

func (c *Client) openAppendOnly(ctx context.Context, name string) (io.WriteCloser, error) {
	switch c.csvSpec.Backend {
	case BackendTypeLocal:
		name := path.Join(c.csvSpec.Directory, name)
		return local.OpenAppendOnly(name, c.csvSpec.MaxFileSize)
	case BackendTypeGCS:
		name := path.Join(c.path, name)
		return gcs.OpenAppendOnly(ctx, c.gcpStorageClient, c.baseDir, name, c.csvSpec.MaxFileSize)
	case BackendTypeS3:
		name := path.Join(c.path, name)
		return s3.OpenAppendOnly(ctx, c.awsUploader, c.baseDir, name, c.csvSpec.MaxFileSize)
	default:
		panic("unknown backend " + c.csvSpec.Backend.String())
	}
}
