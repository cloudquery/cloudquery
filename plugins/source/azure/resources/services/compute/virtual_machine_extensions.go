// Auto generated code - DO NOT EDIT.

package compute

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"
)

func virtualMachineExtensions() *schema.Table {
	return &schema.Table{
		Name:        "azure_compute_virtual_machine_extensions",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute#VirtualMachineExtension`,
		Resolver:    fetchComputeVirtualMachineExtensions,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "compute_virtual_machine_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "force_update_tag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ForceUpdateTag"),
			},
			{
				Name:     "publisher",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Publisher"),
			},
			{
				Name:     "type_handler_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TypeHandlerVersion"),
			},
			{
				Name:     "auto_upgrade_minor_version",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AutoUpgradeMinorVersion"),
			},
			{
				Name:     "enable_automatic_upgrade",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableAutomaticUpgrade"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "instance_view",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("InstanceView"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver(`Type`),
			},
		},
	}
}

func fetchComputeVirtualMachineExtensions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Compute.VirtualMachineExtensions

	virtualMachine := parent.Item.(compute.VirtualMachine)
	resource, err := client.ParseResourceID(*virtualMachine.ID)
	if err != nil {
		return err
	}
	response, err := svc.List(ctx, resource.ResourceGroup, *virtualMachine.Name, "")
	if err != nil {
		return err
	}
	if response.Value == nil {
		return nil
	}
	res <- *response.Value

	return nil
}
