// Code generated by codegen; DO NOT EDIT.

package armelasticsan

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/julienschmidt/httprouter"
)

func createRpOperation() (*client.Services, error) {
	var item armelasticsan.OperationsClientListResponse
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
	svc, err := armelasticsan.NewOperationsClient(client.TestSubscription, &client.MockCreds{}, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: ts.Client(),
		},
	})
	if err != nil {
		return nil, err
	}
	return &client.Services{
		ArmelasticsanRpOperation: svc,
	}, nil
}

func TestRpOperation(t *testing.T) {
	client.MockTestHelper(t, RpOperation(), createRpOperation)
}
