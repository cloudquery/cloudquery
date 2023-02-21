package coupons_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/coupons"
)

func TestCoupons(t *testing.T) {
	client.MockTestHelper(t, coupons.Coupons(), client.TestOptions{})
}
