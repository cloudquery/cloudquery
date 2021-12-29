// +build integration

package mq

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationMQBrokers(t *testing.T) {
	client.AWSTestHelper(t, MqBrokers(),
		"./snapshots")
}
