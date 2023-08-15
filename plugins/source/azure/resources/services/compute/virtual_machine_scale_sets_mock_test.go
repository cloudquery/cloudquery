package compute

import (
	"encoding/json"
	"net/http"

	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createVirtualMachineScaleSets(router *mux.Router) error {
	var item armcompute.VirtualMachineScaleSetsClientListAllResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/virtualMachineScaleSets", func(w http.ResponseWriter, r *http.Request) {
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

	var vm armcompute.VirtualMachineScaleSetVMsClientListResponse
	if err := faker.FakeObject(&vm); err != nil {
		return err
	}
	vm.NextLink = &emptyStr
	router.HandleFunc("/subscriptions/{subscriptionId}/resourceGroups/{group}/providers/Microsoft.Compute/virtualMachineScaleSets/{scaleSet}/virtualMachines", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&vm)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return createVirtualMachineScaleSetsNetworkInterfaces(router)
}

func TestVirtualMachineScaleSets(t *testing.T) {
	client.MockTestHelper(t, VirtualMachineScaleSets(), createVirtualMachineScaleSets)
}
