//go:build integration
// +build integration

package emr

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationEmrClusters(t *testing.T) {
	client.AWSTestHelper(t, EmrClusters())
}
