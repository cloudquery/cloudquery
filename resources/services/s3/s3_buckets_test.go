// +build integration

package s3

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationS3Buckets(t *testing.T) {
	client.AWSTestHelper(t, S3Buckets(),
		"./snapshots")
}
