package reporting_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/reporting"
)

func TestReportingReportRuns(t *testing.T) {
	client.MockTestHelper(t, reporting.ReportingReportRuns(), client.TestOptions{})
}
