package sigma_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/sigma"
)

func TestSigmaScheduledQueryRuns(t *testing.T) {
	client.MockTestHelper(t, sigma.SigmaScheduledQueryRuns(), client.TestOptions{})
}
