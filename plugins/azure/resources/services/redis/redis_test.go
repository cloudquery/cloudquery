//go:build integration

package redis

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationRedisServices(t *testing.T) {
	client.AzureTestHelper(t, RedisServices())
}
