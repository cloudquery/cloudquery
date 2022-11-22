package gcs

import (
	"context"
	"io"
	"io/fs"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
)

type fileWriter struct {
	ctx           context.Context
	storageClient *storage.Client
	writer        *storage.Writer
	written       uint64
	bucket        string
	name          string
	maxFileSize   uint64
}

type fileReader struct {
	ctx           context.Context
	storageClient *storage.Client
	reader        *storage.Reader
	bucket        string
	name          string
}

func (f *fileReader) Read(p []byte) (n int, err error) {
	return f.reader.Read(p)
}

func (f *fileReader) Close() error {
	defer func() {
		f.reader = nil
	}()
	if f.reader == nil {
		return fs.ErrClosed
	}
	return f.reader.Close()
}

func (f *fileWriter) Write(data []byte) (int, error) {
	n, err := f.writer.Write(data)
	if err != nil {
		return n, err
	}
	f.written += uint64(n)
	if f.maxFileSize != 0 && f.written >= f.maxFileSize {
		if err := f.writer.Close(); err != nil {
			return n, err
		}
		f.written = 0
		name := f.name + "." + uuid.NewString()
		f.writer = f.storageClient.Bucket(f.bucket).Object(name).NewWriter(f.ctx)
	}
	return n, nil
}

func (f *fileWriter) Close() error {
	defer func() {
		f.writer = nil
		f.written = 0
	}()
	if f.writer == nil {
		return fs.ErrClosed
	}
	return f.writer.Close()
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
	return &fileWriter{
		ctx:           ctx,
		storageClient: storageClient,
		writer:        storageClient.Bucket(bucket).Object(uniqueName).NewWriter(ctx),
		bucket:        bucket,
		name:          name,
		maxFileSize:   maxFileSize,
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
	return &fileReader{
		ctx:           ctx,
		storageClient: storageClient,
		reader:        reader,
		bucket:        bucket,
		name:          name,
	}, nil
}
