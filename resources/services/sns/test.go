// +build integration

package sns

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationSnsSubscriptions(t *testing.T) {
	client.AWSTestHelper(t, SnsSubscriptions())
}

func TestIntegrationSnsTopics(t *testing.T) {
	client.AWSTestHelper(t, SnsTopics())
}
