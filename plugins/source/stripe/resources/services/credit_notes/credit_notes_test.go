package credit_notes_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/credit_notes"
)

func TestCreditNotes(t *testing.T) {
	client.MockTestHelper(t, credit_notes.CreditNotes(), client.TestOptions{})
}
