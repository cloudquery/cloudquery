package client

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
)

type gcsFile struct {
	ctx 				 context.Context
	storageClient *storage.Client
	gcsWriter *storage.Writer
	written  uint64
	bucket string
	name string
}

func (f *gcsFile) Write(data []byte) (int, error) {
	n, err := f.gcsWriter.Write(data)
	if err != nil {
		return n, err
	}
	f.written += uint64(n)
	if f.written >= batchSize {
		if err := f.gcsWriter.Close(); err != nil {
			return n, err
		}
		f.written = 0
		name := uuid.NewString() + "." + f.name
		f.gcsWriter = f.storageClient.Bucket(f.bucket).Object(name).NewWriter(f.ctx)
	}
	return n, nil
}

func (f *gcsFile) Close() error {
	return f.gcsWriter.Close()
}

func (c *Client) openGCSFileAppend(ctx context.Context, name string) (io.Writer, error) {
	return &gcsFile{
		ctx: ctx,
		storageClient: c.gcpStorageClient,
		gcsWriter: c.gcpStorageClient.Bucket(c.bucket).Object(name).NewWriter(ctx),
		bucket: c.bucket,
		name: name,
	}, nil
}