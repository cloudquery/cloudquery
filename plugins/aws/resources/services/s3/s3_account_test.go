// +build integration

package s3

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationS3Account(t *testing.T) {
	client.AWSTestHelper(t, S3Accounts(),
		"./snapshots")
}
