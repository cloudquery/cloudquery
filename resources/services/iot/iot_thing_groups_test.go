//go:build integration
// +build integration

package iot

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationIotThingGroups(t *testing.T) {
	client.AWSTestHelper(t, IotThingGroups())
}
