package client

import (
	"context"
	"fmt"
	"io"

	"github.com/cloudquery/cloudquery/plugins/destination/file/internal/backends/gcs"
	"github.com/cloudquery/cloudquery/plugins/destination/file/internal/backends/local"
	"github.com/cloudquery/cloudquery/plugins/destination/file/internal/backends/s3"
)


func (c *Client) openReadOnly(name string) (io.Reader, error) {
	// io.Writer
	switch c.csvSpec.Backend {
	case BackendTypeLocal:
		return local.OpenReadOnly(name)
	case BackendTypeGCS:
		return nil, fmt.Errorf("not implemented")
	case BackendTypeS3:
		return nil, fmt.Errorf("not implemented")
	default:
		panic("unknown backend " + c.csvSpec.Backend)
	}
}

func (c *Client) OpenAppendOnly(ctx context.Context, name string) (io.WriteCloser, error) {
	switch c.csvSpec.Backend {
	case BackendTypeLocal:
		return local.OpenAppendOnly(name)
	case BackendTypeGCS:
		return gcs.OpenAppendOnly(ctx, c.gcpStorageClient, c.bucket, name)
	case BackendTypeS3:
		return s3.OpenS3FileAppend(ctx, c.awsUploader, c.bucket, name)
	default:
		panic("unknown backend " + c.csvSpec.Backend)
	}
}

