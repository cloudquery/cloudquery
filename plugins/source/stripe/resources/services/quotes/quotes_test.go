package quotes_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/quotes"
)

func TestQuotes(t *testing.T) {
	client.MockTestHelper(t, quotes.Quotes(), client.TestOptions{})
}
