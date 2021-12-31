// +build integration

package iam

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationIamUsers(t *testing.T) {
	client.AWSTestHelper(t, IamUsers())
}
