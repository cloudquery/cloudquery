// +build integration

package elasticbeanstalk

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationElasticbeanstalkApplications(t *testing.T) {
	client.AWSTestHelper(t, ElasticbeanstalkApplications(),
		"./snapshots")
}
