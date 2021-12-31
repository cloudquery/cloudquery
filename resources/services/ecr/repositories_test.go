// +build integration

package ecr

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationEcrRepositories(t *testing.T) {
	client.AWSTestHelper(t, EcrRepositories())
}
