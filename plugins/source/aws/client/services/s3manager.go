package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//go:generate mockgen -package=mocks -destination=../mocks/s3manager.go . S3managerClient
type S3managerClient interface {
	GetBucketRegion(context.Context, string, ...func(*s3.Options)) (string, error)
}
