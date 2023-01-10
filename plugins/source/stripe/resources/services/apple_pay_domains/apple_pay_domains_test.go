package apple_pay_domains_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/apple_pay_domains"
)

func TestApplePayDomains(t *testing.T) {
	client.MockTestHelper(t, apple_pay_domains.ApplePayDomains(), client.TestOptions{})
}
