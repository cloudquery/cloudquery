//go:build integration
// +build integration

package cloudfront

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationCloudfrontDistributions(t *testing.T) {
	client.AWSTestHelper(t, CloudfrontDistributions())
}
