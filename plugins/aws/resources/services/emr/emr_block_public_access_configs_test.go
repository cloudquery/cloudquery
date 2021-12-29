// +build integration

package emr

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationEmrBlockPublicAccessConfigs(t *testing.T) {
	client.AWSTestHelper(t, EmrBlockPublicAccessConfigs(),
		"./snapshots")
}
