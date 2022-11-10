// Auto generated code - DO NOT EDIT.

package compute

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Disks() *schema.Table {
	return &schema.Table{
		Name:        "azure_compute_disks",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute#Disk`,
		Resolver:    fetchComputeDisks,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "managed_by",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ManagedBy"),
			},
			{
				Name:     "managed_by_extended",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ManagedByExtended"),
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
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
				Name:     "time_created",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("TimeCreated"),
			},
			{
				Name:     "os_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OsType"),
			},
			{
				Name:     "hyper_v_generation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HyperVGeneration"),
			},
			{
				Name:     "purchase_plan",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PurchasePlan"),
			},
			{
				Name:     "creation_data",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CreationData"),
			},
			{
				Name:     "disk_size_gb",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DiskSizeGB"),
			},
			{
				Name:     "disk_size_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DiskSizeBytes"),
			},
			{
				Name:     "unique_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UniqueID"),
			},
			{
				Name:     "encryption_settings_collection",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EncryptionSettingsCollection"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "disk_iops_read_write",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DiskIOPSReadWrite"),
			},
			{
				Name:     "disk_m_bps_read_write",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DiskMBpsReadWrite"),
			},
			{
				Name:     "disk_iops_read_only",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DiskIOPSReadOnly"),
			},
			{
				Name:     "disk_m_bps_read_only",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DiskMBpsReadOnly"),
			},
			{
				Name:     "disk_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskState"),
			},
			{
				Name:     "encryption",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Encryption"),
			},
			{
				Name:     "max_shares",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxShares"),
			},
			{
				Name:     "share_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ShareInfo"),
			},
			{
				Name:     "network_access_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NetworkAccessPolicy"),
			},
			{
				Name:     "disk_access_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskAccessID"),
			},
			{
				Name:     "tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Tier"),
			},
			{
				Name:     "bursting_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("BurstingEnabled"),
			},
			{
				Name:     "property_updates_in_progress",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PropertyUpdatesInProgress"),
			},
			{
				Name:     "supports_hibernation",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SupportsHibernation"),
			},
			{
				Name:     "security_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SecurityProfile"),
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
	}
}

func fetchComputeDisks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Compute.Disks

	response, err := svc.List(ctx)

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
