package reporting_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/reporting"
)

func TestReportingReportTypes(t *testing.T) {
	client.MockTestHelper(t, reporting.ReportingReportTypes(), client.TestOptions{})
}
