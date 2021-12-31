// +build integration

package secretsmanager

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationSecretsmanagerSecrets(t *testing.T) {
	client.AWSTestHelper(t, SecretsmanagerSecrets())
}
