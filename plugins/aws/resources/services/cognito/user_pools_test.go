//go:build integration
// +build integration

package cognito

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationCognitoUserPools(t *testing.T) {
	client.AWSTestHelper(t, CognitoUserPools())
}
