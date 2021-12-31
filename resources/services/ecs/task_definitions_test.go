// +build integration

package ecs

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationEcsTaskDefinitions(t *testing.T) {
	client.AWSTestHelper(t, EcsTaskDefinitions())
}
