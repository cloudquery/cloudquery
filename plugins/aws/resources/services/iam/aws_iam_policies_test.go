// +build integration

package iam

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationIamPolicies(t *testing.T) {
	awsTestIntegrationHelper(t, resources.IamPolicies(),
		"./snapshots")
}
