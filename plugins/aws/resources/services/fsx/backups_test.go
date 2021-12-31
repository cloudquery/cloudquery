// +build integration

package fsx

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationFsxBackups(t *testing.T) {
	client.AWSTestHelper(t, FsxBackups())
}
