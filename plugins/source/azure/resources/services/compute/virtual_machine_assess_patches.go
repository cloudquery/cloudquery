package compute

import (
	"context"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func virtualMachineAssessPatches() *schema.Table {
	tableName := `azure_compute_virtual_machine_assess_patches`
	return &schema.Table{
		Name:        tableName,
		Resolver:    fetchVirtualMachineAssessPatches,
		Description: "https://learn.microsoft.com/en-us/rest/api/compute/virtual-machines/assess-patches?tabs=HTTP#virtualmachineassesspatchesresult",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace(tableName, client.Namespacemicrosoft_compute),
		Transform:   transformers.TransformWithStruct(&armcompute.VirtualMachineAssessPatchesResult{}, transformers.WithPrimaryKeys("AssessmentActivityID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
	}
}

func fetchVirtualMachineAssessPatches(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*armcompute.VirtualMachine)
	cl := meta.(*client.Client)
	svc, err := armcompute.NewVirtualMachinesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*p.ID)
	if err != nil {
		return err
	}
	ap, err := svc.BeginAssessPatches(ctx, group, *p.Name, nil)
	if err != nil {
		return err
	}
	resp, err := ap.PollUntilDone(ctx, &runtime.PollUntilDoneOptions{
		Frequency: 5 * time.Second,
	})
	if err != nil {
		return err
	}
	res <- resp.VirtualMachineAssessPatchesResult
	return nil
}
