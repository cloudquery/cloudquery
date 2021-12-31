// +build integration

package redshift

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationRedshiftSubnetGroups(t *testing.T) {
	client.AWSTestHelper(t, RedshiftSubnetGroups())
}
