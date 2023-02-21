package identity_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/identity"
)

func TestIdentityVerificationReports(t *testing.T) {
	client.MockTestHelper(t, identity.IdentityVerificationReports(), client.TestOptions{})
}
