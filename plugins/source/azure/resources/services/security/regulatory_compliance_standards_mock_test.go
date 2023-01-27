package security

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func createRegulatoryComplianceStandards(router *mux.Router) error {
	var item armsecurity.RegulatoryComplianceStandardsClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&item)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return nil
}

func TestRegulatoryComplianceStandards(t *testing.T) {
	client.MockTestHelper(t, RegulatoryComplianceStandards(), createRegulatoryComplianceStandards)
}

func TestCheckNoStandardPricingBundle(t *testing.T) {
	type testCase struct {
		err      error
		expected error
	}

	for _, tc := range []testCase{
		{
			err:      io.EOF,
			expected: io.EOF,
		},
		{
			err:      &azcore.ResponseError{},
			expected: &azcore.ResponseError{},
		},
		{
			err: &azcore.ResponseError{ErrorCode: `Subscription with no standard pricing bundle`},
		},
	} {
		require.EqualValues(t, tc.expected, checkNoStandardPricingBundle(tc.err))
	}
}
