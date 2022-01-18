//go:build integration
// +build integration

package cognito

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationCognitoIdentityPools(t *testing.T) {
	client.AWSTestHelper(t, CognitoIdentityPools())
}
