// +build integration

package route53

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationRoute53DelegationSet(t *testing.T) {
	client.AWSTestHelper(t, Route53ReusableDelegationSets())
}
