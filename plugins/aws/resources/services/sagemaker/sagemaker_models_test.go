// +build integration

package sagemaker

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationSageMakerModels(t *testing.T) {
	client.AWSTestHelper(t, SagemakerModels(),
		"./snapshots")
}
