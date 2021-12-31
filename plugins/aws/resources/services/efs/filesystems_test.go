// +build integration

package efs

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationEfsFilesystems(t *testing.T) {
	client.AWSTestHelper(t, EfsFilesystems())
}
