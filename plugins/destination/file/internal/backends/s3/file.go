package s3

import (
	"bytes"
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

type file struct {
	ctx         context.Context
	writer      io.Writer
	reader      io.Reader
	uploader    *manager.Uploader
	downloader  *manager.Downloader
	eg          *errgroup.Group
	written     uint64
	bucket      string
	name        string
	maxFileSize uint64
}

func (f *file) Read(p []byte) (n int, err error) {
	return f.reader.Read(p)
}

func (f *file) Write(data []byte) (int, error) {
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
			_, err := f.uploader.Upload(f.ctx, &s3.PutObjectInput{
				Bucket: aws.String(f.bucket),
				Key:    aws.String(uniqueName),
				Body:   f.reader,
			})
			return err
		})
	}
	return n, nil
}

func (f *file) Close() error {
	if f.eg != nil {
		if err := f.writer.(*io.PipeWriter).Close(); err != nil {
			return err
		}
		return f.eg.Wait()
	}
	return nil
}

func OpenAppendOnly(
	ctx context.Context,
	uploader *manager.Uploader,
	bucket string,
	name string,
	maxFileSize uint64) (io.WriteCloser, error) {
	uniqueName := name
	if maxFileSize != 0 {
		uniqueName = name + "." + uuid.NewString()
	}
	r, w := io.Pipe()
	f := file{
		ctx:      ctx,
		bucket:   bucket,
		name:     uniqueName,
		uploader: uploader,
		writer:   w,
		reader:   r,
		eg:       &errgroup.Group{},
	}
	f.eg.Go(func() error {
		_, err := uploader.Upload(ctx, &s3.PutObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(uniqueName),
			Body:   r,
		})
		return err
	})
	return &f, nil
}

func OpenReadOnly(
	ctx context.Context,
	downloader *manager.Downloader,
	bucket string,
	name string) (io.ReadCloser, error) {
	f := file{
		ctx:        ctx,
		bucket:     bucket,
		name:       name,
		downloader: downloader,
	}
	// we are downloading everything into memory because we only
	// using it for testing and implementing WriterAt and Reader interface
	// is quite tricky so skipping this for now.
	writerAtBuffer := manager.NewWriteAtBuffer(make([]byte, 0, 1024*1024))
	_, err := downloader.Download(ctx,
		writerAtBuffer,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(name),
		})
	if err != nil {
		return nil, err
	}
	f.reader = bytes.NewReader(writerAtBuffer.Bytes())
	return &f, nil
}
