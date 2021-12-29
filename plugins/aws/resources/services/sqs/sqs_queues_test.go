// +build integration

package sqs

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationSQSQueues(t *testing.T) {
	client.AWSTestHelper(t, SQSQueues(),
		"./snapshots")
}
