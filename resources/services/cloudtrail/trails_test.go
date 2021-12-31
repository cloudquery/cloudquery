// +build integration

package cloudtrail

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationCloudtrailTrails(t *testing.T) {
	client.AWSTestHelper(t, CloudtrailTrails())
}
