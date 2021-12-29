// +build integration

package cognito

import (
	"github.com/cloudquery/cq-provider-aws/client"
	"testing"
)

func TestIntegrationCognitoUserPools(t *testing.T) {
	client.AWSTestHelper(t, CognitoUserPools(),
		"./snapshots")
}
