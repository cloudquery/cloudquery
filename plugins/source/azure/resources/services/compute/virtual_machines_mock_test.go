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

func createVirtualMachines(router *mux.Router) error {
	var item armcompute.VirtualMachinesClientListAllResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr
	s := armcompute.WindowsVMGuestPatchModeAutomaticByPlatform
	for i := range item.VirtualMachineListResult.Value {
		item.VirtualMachineListResult.Value[i].Properties.OSProfile.WindowsConfiguration.PatchSettings = &armcompute.PatchSettings{
			PatchMode: &s, // required for virtual_machine_patch_assessments_mock_test.go
		}
	}

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/virtualMachines", func(w http.ResponseWriter, r *http.Request) {
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

	var instanceViewItem armcompute.VirtualMachinesClientInstanceViewResponse
	if err := faker.FakeObject(&instanceViewItem); err != nil {
		return err
	}
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/instanceView
	router.HandleFunc("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/instanceView", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&instanceViewItem)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	if err := createVirtualMachineAssessPatches(router); err != nil {
		return err
	}

	return createVirtualMachineExtensions(router)
}

func TestVirtualMachines(t *testing.T) {
	client.MockTestHelper(t, VirtualMachines(), createVirtualMachines)
}
