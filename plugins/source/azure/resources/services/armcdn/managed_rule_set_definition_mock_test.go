// Code generated by codegen; DO NOT EDIT.

package armcdn

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/julienschmidt/httprouter"
)

func createManagedRuleSetDefinition() (*client.Services, error) {
	var item armcdn.ManagedRuleSetsClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return nil, err
	}
	emptyStr := ""
	item.NextLink = &emptyStr
	mux := httprouter.New()
	mux.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	ts := httptest.NewServer(mux)
	cloud.AzurePublic.Services[cloud.ResourceManager] = cloud.ServiceConfiguration{
		Endpoint: ts.URL,
		Audience: "test",
	}
	svc, err := armcdn.NewManagedRuleSetsClient(client.TestSubscription, &client.MockCreds{}, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: ts.Client(),
		},
	})
	if err != nil {
		return nil, err
	}
	return &client.Services{
		ArmcdnManagedRuleSetDefinition: svc,
	}, nil
}

func TestManagedRuleSetDefinition(t *testing.T) {
	client.MockTestHelper(t, ManagedRuleSetDefinition(), createManagedRuleSetDefinition)
}
