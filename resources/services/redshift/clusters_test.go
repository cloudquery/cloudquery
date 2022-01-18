//go:build integration
// +build integration

package redshift

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationRedshiftClusters(t *testing.T) {
	client.AWSTestHelper(t, RedshiftClusters())
}
