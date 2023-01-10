package reviews_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/reviews"
)

func TestReviews(t *testing.T) {
	client.MockTestHelper(t, reviews.Reviews(), client.TestOptions{})
}
