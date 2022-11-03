package s3manager

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	s3Client *s3.Client
}

func NewFromConfig(cfg aws.Config) Client {
	return Client{
		s3Client: s3.NewFromConfig(cfg),
	}
}

func (c Client) GetBucketRegion(ctx context.Context, bucket string, optFns ...func(*s3.Options)) (string, error) {
	return manager.GetBucketRegion(ctx, c.s3Client, bucket, optFns...)
}
