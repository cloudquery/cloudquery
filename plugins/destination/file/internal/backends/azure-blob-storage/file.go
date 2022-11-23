package azure_blob_storage

import (
	"context"
	"io"
	"io/fs"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

type fileWriter struct {
	ctx         context.Context
	writer      io.Writer
	reader      io.Reader
	client      *azblob.Client
	eg          *errgroup.Group
	written     uint64
	container   string
	name        string
	maxFileSize uint64
}

type fileReader struct {
	ctx       context.Context
	reader    io.Reader
	client    *azblob.Client
	container string
	name      string
}

func (f *fileReader) Read(p []byte) (n int, err error) {
	return f.reader.Read(p)
}

func (f *fileReader) Close() error {
	f.reader = nil
	return nil
}

func (f *fileWriter) Write(data []byte) (int, error) {
	n, err := f.writer.Write(data)
	if err != nil {
		return n, err
	}
	f.written += uint64(n)
	if f.maxFileSize != 0 && f.written >= f.maxFileSize {
		if err := f.eg.Wait(); err != nil {
			return n, err
		}
		f.written = 0
		uniqueName := f.name + "." + uuid.NewString()
		f.eg.Go(func() error {
			_, err := f.client.UploadStream(f.ctx, f.container, uniqueName, f.reader, nil)
			return err
		})
	}
	return n, nil
}

func (f *fileWriter) Close() error {
	defer func() {
		f.writer = nil
		f.eg = nil
	}()
	if f.eg == nil || f.writer == nil {
		return fs.ErrClosed
	}
	if err := f.writer.(*io.PipeWriter).Close(); err != nil {
		return err
	}
	return f.eg.Wait()
}

func OpenAppendOnly(
	ctx context.Context,
	client *azblob.Client,
	container string,
	blobName string,
	maxFileSize uint64) (io.WriteCloser, error) {
	uniqueName := blobName
	if maxFileSize != 0 {
		uniqueName = blobName + "." + uuid.NewString()
	}
	r, w := io.Pipe()
	f := fileWriter{
		ctx:       ctx,
		container: container,
		name:      uniqueName,
		client:    client,
		writer:    w,
		reader:    r,
		eg:        &errgroup.Group{},
	}
	f.eg.Go(func() error {
		_, err := f.client.UploadStream(f.ctx, f.container, uniqueName, f.reader, nil)
		return err
	})
	return &f, nil
}

func OpenReadOnly(
	ctx context.Context,
	client *azblob.Client,
	container string,
	blobName string) (io.ReadCloser, error) {
	f := fileReader{
		ctx:       ctx,
		container: container,
		name:      blobName,
		client:    client,
	}
	s, err := client.DownloadStream(ctx,
		container,
		blobName,
		nil,
	)
	if err != nil {
		return nil, err
	}
	f.reader = s.Body
	return &f, nil
}
