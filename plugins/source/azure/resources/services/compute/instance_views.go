// Auto generated code - DO NOT EDIT.

package compute

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"
)

func instanceViews() *schema.Table {
	return &schema.Table{
		Name:     "azure_compute_instance_views",
		Resolver: fetchComputeInstanceViews,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "cq_id_parent",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIdResolver,
			},
			{
				Name:     "platform_update_domain",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PlatformUpdateDomain"),
			},
			{
				Name:     "platform_fault_domain",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PlatformFaultDomain"),
			},
			{
				Name:     "computer_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ComputerName"),
			},
			{
				Name:     "os_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OsName"),
			},
			{
				Name:     "os_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OsVersion"),
			},
			{
				Name:     "hyper_v_generation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HyperVGeneration"),
			},
			{
				Name:     "rdp_thumb_print",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RdpThumbPrint"),
			},
			{
				Name:     "vm_agent",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VMAgent"),
			},
			{
				Name:     "maintenance_redeploy_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MaintenanceRedeployStatus"),
			},
			{
				Name:     "disks",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Disks"),
			},
			{
				Name:     "extensions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Extensions"),
			},
			{
				Name:     "vm_health",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VMHealth"),
			},
			{
				Name:     "boot_diagnostics",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BootDiagnostics"),
			},
			{
				Name:     "assigned_host",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AssignedHost"),
			},
			{
				Name:     "statuses",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Statuses"),
			},
			{
				Name:     "patch_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PatchStatus"),
			},
		},
	}
}

func fetchComputeInstanceViews(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Compute.InstanceViews

	virtualMachine := parent.Item.(compute.VirtualMachine)
	resource, err := client.ParseResourceID(*virtualMachine.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	response, err := svc.InstanceView(ctx, resource.ResourceGroup, *virtualMachine.Name)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- response
	return nil
}
