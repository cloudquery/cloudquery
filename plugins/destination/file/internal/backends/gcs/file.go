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
	gcsReader *storage.Reader
	written  uint64
	bucket string
	name string
	maxFileSize uint64
}

func (f *file) Write(data []byte) (int, error) {
	n, err := f.gcsWriter.Write(data)
	if err != nil {
		return n, err
	}
	f.written += uint64(n)
	if f.maxFileSize != 0 && f.written >= f.maxFileSize {
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
	if f.gcsWriter != nil {
		return f.gcsWriter.Close()
	}
	if f.gcsReader != nil {
		return f.gcsReader.Close()
	}
	return nil
}

func (f *file) Read(p []byte) (n int, err error) {
	return f.gcsReader.Read(p)
}

// OpenAppendOnly opens a file on GCS for writing in append only mode.
// if maxFileSize is 0 then all data will be written to a single file
// otherwise a new file will be created when the maxFileSize is reached (with uuid suffix)
func OpenAppendOnly(
	ctx context.Context,
	storageClient *storage.Client,
	bucket string,
	name string,
	maxFileSize uint64,
) (io.WriteCloser, error) {
	uniqueName := name
	// if maxFileSize is non zero then we need to append a unique name to the file
	// as we need to split the file into chunks
	if maxFileSize != 0 {
		uniqueName = name + "." + uuid.NewString()
	}
	return &file{
		ctx: ctx,
		storageClient: storageClient,
		gcsWriter: storageClient.Bucket(bucket).Object(uniqueName).NewWriter(ctx),
		bucket: bucket,
		name: name,
		maxFileSize: maxFileSize,
	}, nil
}

// OpenReadOnly opens a file on GCS for read only mode.
// if maxFileSize is 0 then all data will be written to a single file
// otherwise a new file will be created when the maxFileSize is reached (with uuid suffix)
func OpenReadOnly(
	ctx context.Context,
	storageClient *storage.Client,
	bucket string,
	name string) (io.ReadCloser, error) {
	reader, err := storageClient.Bucket(bucket).Object(name).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	return &file{
		ctx: ctx,
		storageClient: storageClient,
		gcsReader: reader,
		bucket: bucket,
		name: name,
	}, nil
}