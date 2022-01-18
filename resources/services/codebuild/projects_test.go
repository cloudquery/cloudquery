//go:build integration
// +build integration

package codebuild

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationCodebuildProjects(t *testing.T) {
	client.AWSTestHelper(t, CodebuildProjects())
}
