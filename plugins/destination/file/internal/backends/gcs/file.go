package gcs

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
)

type file struct {
	ctx 				 context.Context
	storageClient *storage.Client
	gcsWriter *storage.Writer
	written  uint64
	bucket string
	name string
}

func (f *file) Write(data []byte) (int, error) {
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
		name := f.name + "." + uuid.NewString()
		f.gcsWriter = f.storageClient.Bucket(f.bucket).Object(name).NewWriter(f.ctx)
	}
	return n, nil
}

func (f *file) Close() error {
	return f.gcsWriter.Close()
}


func OpenAppendOnly(ctx context.Context, storageClient *storage.Client, bucket string, name string) (io.WriteCloser, error) {
	return &file{
		ctx: ctx,
		storageClient: storageClient,
		gcsWriter: storageClient.Bucket(bucket).Object(name).NewWriter(ctx),
		bucket: bucket,
		name: name,
	}, nil
}