// +build integration

package sagemaker

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationSageMakerEndpointConfigurations(t *testing.T) {
	client.AWSTestHelper(t, SagemakerEndpointConfigurations(),
		"./snapshots")
}
