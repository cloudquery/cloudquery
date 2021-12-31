// +build integration

package iam

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIamPasswordPoliciesIntegration(t *testing.T) {
	client.AWSTestHelper(t, IamPasswordPolicies())
}
