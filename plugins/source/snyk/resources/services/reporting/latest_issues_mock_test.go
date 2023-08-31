package reporting

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
)

func TestLatestIssues(t *testing.T) {
	client.MockTestHelper(t, LatestIssues(), createIssuesForPath("/v1/reporting/issues/latest"))
}
