// Auto generated code - DO NOT EDIT.

package compute

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func VirtualMachineScaleSets() *schema.Table {
	return &schema.Table{
		Name:        "azure_compute_virtual_machine_scale_sets",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute#VirtualMachineScaleSet`,
		Resolver:    fetchComputeVirtualMachineScaleSets,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
			},
			{
				Name:     "plan",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Plan"),
			},
			{
				Name:     "upgrade_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("UpgradePolicy"),
			},
			{
				Name:     "automatic_repairs_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AutomaticRepairsPolicy"),
			},
			{
				Name:     "virtual_machine_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VirtualMachineProfile"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "overprovision",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Overprovision"),
			},
			{
				Name:     "unique_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UniqueID"),
			},
			{
				Name:     "single_placement_group",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SinglePlacementGroup"),
			},
			{
				Name:     "zone_balance",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ZoneBalance"),
			},
			{
				Name:     "platform_fault_domain_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PlatformFaultDomainCount"),
			},
			{
				Name:     "proximity_placement_group",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ProximityPlacementGroup"),
			},
			{
				Name:     "host_group",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HostGroup"),
			},
			{
				Name:     "additional_capabilities",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AdditionalCapabilities"),
			},
			{
				Name:     "scale_in_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ScaleInPolicy"),
			},
			{
				Name:     "orchestration_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OrchestrationMode"),
			},
			{
				Name:     "identity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Identity"),
			},
			{
				Name:     "zones",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Zones"),
			},
			{
				Name:     "extended_location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ExtendedLocation"),
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
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
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
				Name:     "do_not_run_extensions_on_overprovisioned_vms",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver(`DoNotRunExtensionsOnOverprovisionedVMs`),
			},
		},
	}
}

func fetchComputeVirtualMachineScaleSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Compute.VirtualMachineScaleSets

	response, err := svc.ListAll(ctx)

	if err != nil {
		return err
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}

	return nil
}
