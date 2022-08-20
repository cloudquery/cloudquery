package cloudbilling

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/cloudbilling/v1"
	"google.golang.org/api/option"
)

func createAccountsTestServer() (*client.Services, error) {
	ctx := context.Background()
	var account cloudbilling.BillingAccount
	if err := faker.FakeData(&account); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/v1/billingAccounts", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &cloudbilling.ListBillingAccountsResponse{
			BillingAccounts: []*cloudbilling.BillingAccount{&account},
		}
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	var bi cloudbilling.ProjectBillingInfo
	if err := faker.FakeData(&bi); err != nil {
		return nil, err
	}
	bi.BillingAccountName = account.Name
	bi.ProjectId = "testProject"
	mux.GET(fmt.Sprintf("/v1/%s/projects", account.Name), func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &cloudbilling.ListProjectBillingInfoResponse{
			ProjectBillingInfo: []*cloudbilling.ProjectBillingInfo{&bi},
		}
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	ts := httptest.NewServer(mux)
	svc, err := cloudbilling.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		CloudBilling: svc,
	}, nil
}

func TestAccounts(t *testing.T) {
	client.GcpMockTestHelper(t, Accounts(), createAccountsTestServer, client.TestOptions{})
}
