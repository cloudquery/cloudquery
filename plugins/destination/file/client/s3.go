package client

import (
	"context"
	"io"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type s3File struct {
	ctx 				 context.Context
	writer io.Writer
	uploader *manager.Uploader
	wg sync.WaitGroup
	written  uint64
	bucket string
	name string
}

func (f *s3File) Write(data []byte) (int, error) {
	n, err := f.writer.Write(data)
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

func (f *s3File) Close() error {
	return f.gcsWriter.Close()
}

func (c *Client) openS3FileAppend(ctx context.Context, name string) (io.Writer, error) {
	uniqueName := uuid.NewString() + "." + name
	r, w := io.Pipe()
	f := s3File{
		ctx: ctx,
		bucket: c.bucket,
		name: name,
		uploader: c.awsUploader,
		writer: w,
	}
	f.wg.Add(1)
	go func() {
		defer f.wg.Done()
		if _, err := c.awsUploader.Upload(ctx, &s3.PutObjectInput{
			Bucket: aws.String(c.bucket),
			Key:    aws.String(uniqueName),
			Body:   r,
		}); err != nil {
			return err
		}
	}()
	f.uploader = c.awsUploader
	return &f, nil
}