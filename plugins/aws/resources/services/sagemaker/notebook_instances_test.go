//go:build integration
// +build integration

package sagemaker

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationSageMakerNotebookInstances(t *testing.T) {
	client.AWSTestHelper(t, SagemakerNotebookInstances())
}
