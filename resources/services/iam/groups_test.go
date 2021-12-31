// +build integration

package iam

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationIamGroups(t *testing.T) {
	client.AWSTestHelper(t, IamGroups())
}
