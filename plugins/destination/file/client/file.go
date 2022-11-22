package client

import (
	"context"
	"io"

	"github.com/cloudquery/cloudquery/plugins/destination/file/internal/backends/gcs"
	"github.com/cloudquery/cloudquery/plugins/destination/file/internal/backends/local"
	"github.com/cloudquery/cloudquery/plugins/destination/file/internal/backends/s3"
)


func (c *Client) openReadOnly(ctx context.Context, name string) (io.ReadCloser, error) {
	// io.Writer
	switch c.csvSpec.Backend {
	case BackendTypeLocal:
		return local.OpenReadOnly(name)
	case BackendTypeGCS:
		return gcs.OpenReadOnly(ctx, c.gcpStorageClient, c.bucket, name)
	case BackendTypeS3:
		return s3.OpenReadOnly(ctx, c.awsDownloader, c.bucket, name)
	default:
		panic("unknown backend " + c.csvSpec.Backend)
	}
}

func (c *Client) openAppendOnly(ctx context.Context, name string) (io.WriteCloser, error) {
	switch c.csvSpec.Backend {
	case BackendTypeLocal:
		return local.OpenAppendOnly(name, c.csvSpec.MaxFileSize)
	case BackendTypeGCS:
		return gcs.OpenAppendOnly(ctx, c.gcpStorageClient, c.bucket, name, c.csvSpec.MaxFileSize)
	case BackendTypeS3:
		return s3.OpenAppendOnly(ctx, c.awsUploader, c.bucket, name, c.csvSpec.MaxFileSize)
	default:
		panic("unknown backend " + c.csvSpec.Backend)
	}
}

