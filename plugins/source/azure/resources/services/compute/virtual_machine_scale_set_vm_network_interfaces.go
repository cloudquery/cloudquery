package compute

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/gorilla/mux"
)

func virtualMachineScaleSetsNetworkInterfaces() *schema.Table {
	return &schema.Table{
		Name:        "azure_compute_virtual_machine_scale_set_network_interfaces",
		Resolver:    fetchVirtualMachineScaleSetsNetworkInterfaces,
		Description: "https://learn.microsoft.com/en-us/rest/api/virtualnetwork/network-interface-in-vm-ss/list-virtual-machine-scale-set-network-interfaces?tabs=HTTP#networkinterface",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_compute_virtual_machine_scale_set_network_interfaces", client.Namespacemicrosoft_network),
		Transform:   transformers.TransformWithStruct(&armnetwork.Interface{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
	}
}

func fetchVirtualMachineScaleSetsNetworkInterfaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	scaleSet := parent.Item.(*armcompute.VirtualMachineScaleSet)
	svc, err := armnetwork.NewInterfacesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*scaleSet.ID)
	if err != nil {
		return err
	}
	pager := svc.NewListVirtualMachineScaleSetNetworkInterfacesPager(group, *scaleSet.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

func createVirtualMachineScaleSetsNetworkInterfaces(router *mux.Router) error {
	var item armnetwork.InterfacesClientListVirtualMachineScaleSetNetworkInterfacesResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/networkInterfaces", func(w http.ResponseWriter, r *http.Request) {
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
