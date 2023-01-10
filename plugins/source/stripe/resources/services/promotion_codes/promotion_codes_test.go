package promotion_codes_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/promotion_codes"
)

func TestPromotionCodes(t *testing.T) {
	client.MockTestHelper(t, promotion_codes.PromotionCodes(), client.TestOptions{})
}
