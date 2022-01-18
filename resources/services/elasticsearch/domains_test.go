//go:build integration
// +build integration

package elasticsearch

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationElasticsearchDomains(t *testing.T) {
	client.AWSTestHelper(t, ElasticsearchDomains())
}
