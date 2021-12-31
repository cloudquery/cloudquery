// +build integration

package rds

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationRdsClusterParameterGroups(t *testing.T) {
	client.AWSTestHelper(t, RdsClusterParameterGroups())
}
