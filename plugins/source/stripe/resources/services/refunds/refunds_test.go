// Code generated by codegen; DO NOT EDIT.

package refunds_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/refunds"
)

func TestRefundsRefund(t *testing.T) {
	client.MockTestHelper(t, refunds.Refunds())
}
