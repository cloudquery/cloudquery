//go:build integration
// +build integration

package iot

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationIotCertificates(t *testing.T) {
	client.AWSTestHelper(t, IotCertificates())
}
