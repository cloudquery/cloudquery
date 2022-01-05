//go:build integration
// +build integration

package bigquery

import (
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
)

func TestIntegrationBigqueryDatasets(t *testing.T) {
	client.GcpTestHelper(t, BigqueryDatasets())
}
