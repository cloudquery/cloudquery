// +build integration

package iam

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationIamOpenidConnectIdentityProviders(t *testing.T) {
	client.AWSTestHelper(t, IamOpenidConnectIdentityProviders())
}
