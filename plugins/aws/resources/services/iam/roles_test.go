// +build integration

package iam

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationIamRoles(t *testing.T) {
	client.AWSTestHelper(t, IamRoles())
}
