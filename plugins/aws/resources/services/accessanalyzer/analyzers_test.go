// +build integration

package accessanalyzer

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationAccessAnalyzers(t *testing.T) {
	client.AWSTestHelper(t, AccessAnalyzerAnalyzer())
}
