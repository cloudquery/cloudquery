// +build integration

package cloudfront

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationCloudfrontCachePolicies(t *testing.T) {
	client.AWSTestHelper(t, CloudfrontCachePolicies(),
		"./snapshots")
}
