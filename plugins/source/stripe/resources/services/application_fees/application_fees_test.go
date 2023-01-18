package application_fees_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/application_fees"
)

func TestApplicationFees(t *testing.T) {
	client.MockTestHelper(t, application_fees.ApplicationFees(), client.TestOptions{})
}
