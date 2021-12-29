// +build integration

package codebuild

import (
	"github.com/cloudquery/cq-provider-aws/client"
	"testing"
)

func TestIntegrationCodebuildProjects(t *testing.T) {
	client.AWSTestHelper(t, CodebuildProjects(),
		"./snapshots")
}
