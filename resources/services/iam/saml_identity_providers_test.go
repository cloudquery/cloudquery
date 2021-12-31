// +build integration

package iam

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationIamSAMLIdentityProviders(t *testing.T) {
	client.AWSTestHelper(t, IamSamlIdentityProviders())
}
