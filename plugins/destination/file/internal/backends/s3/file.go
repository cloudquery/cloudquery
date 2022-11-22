package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

type file struct {
	ctx 				 context.Context
	writer io.Writer
	reader io.Reader
	uploader *manager.Uploader
	eg *errgroup.Group
	written  uint64
	bucket string
	name string
}

func (f *file) Write(data []byte) (int, error) {
	n, err := f.writer.Write(data)
	if err != nil {
		return n, err
	}
	f.written += uint64(n)
	if f.written >= batchSize {
		if err := f.eg.Wait(); err != nil {
			return n, err
		}
		f.written = 0
		uniqueName := uuid.NewString() + "." + f.name
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
	return f.eg.Wait()
}

func OpenS3FileAppend(ctx context.Context, uploader *manager.Uploader, bucket string, name string) (io.WriteCloser, error) {
	uniqueName := name + "." + uuid.NewString()
	r, w := io.Pipe()
	f := file{
		ctx: ctx,
		bucket: bucket,
		name: name,
		uploader: uploader,
		writer: w,
		reader: r,
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