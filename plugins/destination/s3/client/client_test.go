package client

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

func getClient(t *testing.T) *s3.Client {
	t.Helper()

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	return s3.NewFromConfig(cfg)
}

func ensureBucketExists(t *testing.T, bucket string) {
	t.Helper()
	s3Client := getClient(t)
	_, err := s3Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		t.Fatalf("failed to create bucket %s: %v", bucket, err)
	}
}

func TestPluginCSV(t *testing.T) {
	const bucket = "cq-playground-test-csv"
	ensureBucketExists(t, bucket)

	p := destination.NewPlugin("s3", "development", New, destination.WithManagedWriter())

	destination.PluginTestSuiteRunner(t, p,
		Spec{
			Bucket:   bucket,
			Path:     t.TempDir(),
			Format:   FormatTypeCSV,
			NoRotate: true,
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:    true,
			SkipDeleteStale:  true,
			SkipSecondAppend: true,
		},
	)
}

func TestPluginJSON(t *testing.T) {
	const bucket = "cq-playground-test-json"
	ensureBucketExists(t, bucket)

	p := destination.NewPlugin("s3", "development", New, destination.WithManagedWriter())

	destination.PluginTestSuiteRunner(t, p,
		Spec{
			Bucket:   bucket,
			Path:     t.TempDir(),
			Format:   FormatTypeJSON,
			NoRotate: true,
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:    true,
			SkipDeleteStale:  true,
			SkipSecondAppend: true,
		},
	)
}
