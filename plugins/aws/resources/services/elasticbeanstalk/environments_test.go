//go:build integration
// +build integration

package elasticbeanstalk

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationElasticbeanstalkEnvironments(t *testing.T) {
	client.AWSTestHelper(t, ElasticbeanstalkEnvironments())
}
