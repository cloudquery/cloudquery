package compute

import (
	"context"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func virtualMachinePatchAssessments() *schema.Table {
	tableName := `azure_compute_virtual_machine_patch_assessments`
	return &schema.Table{
		Name:                 tableName,
		Resolver:             fetchVirtualMachinePatchAssessments,
		PostResourceResolver: client.LowercaseIDResolver,
		Description: `https://learn.microsoft.com/en-us/rest/api/compute/virtual-machines/assess-patches?tabs=HTTP#virtualmachineassesspatchesresult.

This will begin patch assessments on available virtual machines and can take long to complete.

Not available for all VMs. More at https://learn.microsoft.com/en-us/azure/virtual-machines/automatic-vm-guest-patching#requirements-for-enabling-automatic-vm-guest-patching
`,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace(tableName, client.Namespacemicrosoft_compute),
		Transform: transformers.TransformWithStruct(&armcompute.VirtualMachineAssessPatchesResult{}, transformers.WithPrimaryKeys("AssessmentActivityID")),
		Columns:   schema.ColumnList{client.SubscriptionID},
	}
}

func fetchVirtualMachinePatchAssessments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*armcompute.VirtualMachine)

	supported := false
	if p.Properties != nil && p.Properties.OSProfile != nil {
		if p.Properties.OSProfile.WindowsConfiguration != nil &&
			p.Properties.OSProfile.WindowsConfiguration.PatchSettings != nil &&
			p.Properties.OSProfile.WindowsConfiguration.PatchSettings.PatchMode != nil &&
			*p.Properties.OSProfile.WindowsConfiguration.PatchSettings.PatchMode == armcompute.WindowsVMGuestPatchModeAutomaticByPlatform {
			supported = true
		}
		if p.Properties.OSProfile.LinuxConfiguration != nil &&
			p.Properties.OSProfile.LinuxConfiguration.PatchSettings != nil &&
			p.Properties.OSProfile.LinuxConfiguration.PatchSettings.PatchMode != nil &&
			*p.Properties.OSProfile.LinuxConfiguration.PatchSettings.PatchMode == armcompute.LinuxVMGuestPatchModeAutomaticByPlatform {
			supported = true
		}
	}
	if !supported {
		return nil
	}

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
