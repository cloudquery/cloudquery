// Code generated by codegen; DO NOT EDIT.

package armcompute

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/julienschmidt/httprouter"
)

func createDiskAccess() (*arm.ClientOptions, error) {
	var item armcompute.DiskAccessesClientListResponse
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
	return &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: ts.Client(),
		},
	}, nil
}

func TestDiskAccess(t *testing.T) {
	client.MockTestHelper(t, DiskAccess(), createDiskAccess)
}
