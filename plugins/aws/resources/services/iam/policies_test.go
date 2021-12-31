// +build integration

package iam

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationIamPolicies(t *testing.T) {
	client.AWSTestHelper(t, IamPolicies())
}
