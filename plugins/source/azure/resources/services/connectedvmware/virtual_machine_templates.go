package connectedvmware

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func VirtualMachineTemplates() *schema.Table {
	return &schema.Table{
		Name:                 "azure_connectedvmware_virtual_machine_templates",
		Resolver:             fetchVirtualMachineTemplates,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware@v0.1.0#VirtualMachineTemplate",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_connectedvmware_virtual_machine_templates", client.Namespacemicrosoft_connectedvmwarevsphere),
		Transform:            transformers.TransformWithStruct(&armconnectedvmware.VirtualMachineTemplate{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchVirtualMachineTemplates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconnectedvmware.NewVirtualMachineTemplatesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
