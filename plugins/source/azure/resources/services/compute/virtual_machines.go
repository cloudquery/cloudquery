// Auto generated code - DO NOT EDIT.

package compute

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"
)

func VirtualMachines() *schema.Table {
	return &schema.Table{
		Name:      "azure_compute_virtual_machines",
		Resolver:  fetchComputeVirtualMachines,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "plan",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Plan"),
			},
			{
				Name:     "hardware_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HardwareProfile"),
			},
			{
				Name:     "storage_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StorageProfile"),
			},
			{
				Name:     "additional_capabilities",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AdditionalCapabilities"),
			},
			{
				Name:     "os_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OsProfile"),
			},
			{
				Name:     "network_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkProfile"),
			},
			{
				Name:     "security_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SecurityProfile"),
			},
			{
				Name:     "diagnostics_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DiagnosticsProfile"),
			},
			{
				Name:     "availability_set",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AvailabilitySet"),
			},
			{
				Name:     "virtual_machine_scale_set",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VirtualMachineScaleSet"),
			},
			{
				Name:     "proximity_placement_group",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ProximityPlacementGroup"),
			},
			{
				Name:     "priority",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Priority"),
			},
			{
				Name:     "eviction_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EvictionPolicy"),
			},
			{
				Name:     "billing_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BillingProfile"),
			},
			{
				Name:     "host",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Host"),
			},
			{
				Name:     "host_group",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HostGroup"),
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
				Name:     "license_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LicenseType"),
			},
			{
				Name:     "vmid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VMID"),
			},
			{
				Name:     "extensions_time_budget",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExtensionsTimeBudget"),
			},
			{
				Name:     "platform_fault_domain",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PlatformFaultDomain"),
			},
			{
				Name:     "scheduled_events_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ScheduledEventsProfile"),
			},
			{
				Name:     "user_data",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserData"),
			},
			{
				Name:     "resources",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Resources"),
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
		},

		Relations: []*schema.Table{
			instanceViews(), virtualMachineExtensions(),
		},
	}
}

func fetchComputeVirtualMachines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Compute.VirtualMachines

	response, err := svc.ListAll(ctx, "false")

	if err != nil {
		return errors.WithStack(err)
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
