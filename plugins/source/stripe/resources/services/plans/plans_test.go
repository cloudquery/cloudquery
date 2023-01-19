package plans_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/plans"
)

func TestPlans(t *testing.T) {
	client.MockTestHelper(t, plans.Plans(), client.TestOptions{})
}
