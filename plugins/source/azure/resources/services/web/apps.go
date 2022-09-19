// Auto generated code - DO NOT EDIT.

package web

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Apps() *schema.Table {
	return &schema.Table{
		Name:      "azure_web_apps",
		Resolver:  fetchWebApps,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "host_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("HostNames"),
			},
			{
				Name:     "repository_site_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RepositorySiteName"),
			},
			{
				Name:     "usage_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UsageState"),
			},
			{
				Name:     "enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Enabled"),
			},
			{
				Name:     "enabled_host_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("EnabledHostNames"),
			},
			{
				Name:     "availability_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvailabilityState"),
			},
			{
				Name:     "host_name_ssl_states",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HostNameSslStates"),
			},
			{
				Name:     "server_farm_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerFarmID"),
			},
			{
				Name:     "reserved",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Reserved"),
			},
			{
				Name:     "is_xenon",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsXenon"),
			},
			{
				Name:     "hyper_v",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("HyperV"),
			},
			{
				Name:     "last_modified_time_utc",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LastModifiedTimeUtc"),
			},
			{
				Name:     "site_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SiteConfig"),
			},
			{
				Name:     "traffic_manager_host_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("TrafficManagerHostNames"),
			},
			{
				Name:     "scm_site_also_stopped",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ScmSiteAlsoStopped"),
			},
			{
				Name:     "target_swap_slot",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TargetSwapSlot"),
			},
			{
				Name:     "hosting_environment_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HostingEnvironmentProfile"),
			},
			{
				Name:     "client_affinity_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ClientAffinityEnabled"),
			},
			{
				Name:     "client_cert_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ClientCertEnabled"),
			},
			{
				Name:     "client_cert_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClientCertMode"),
			},
			{
				Name:     "client_cert_exclusion_paths",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClientCertExclusionPaths"),
			},
			{
				Name:     "host_names_disabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("HostNamesDisabled"),
			},
			{
				Name:     "custom_domain_verification_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomDomainVerificationID"),
			},
			{
				Name:     "outbound_ip_addresses",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OutboundIPAddresses"),
			},
			{
				Name:     "possible_outbound_ip_addresses",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PossibleOutboundIPAddresses"),
			},
			{
				Name:     "container_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ContainerSize"),
			},
			{
				Name:     "daily_memory_time_quota",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DailyMemoryTimeQuota"),
			},
			{
				Name:     "suspended_till",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SuspendedTill"),
			},
			{
				Name:     "max_number_of_workers",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxNumberOfWorkers"),
			},
			{
				Name:     "cloning_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CloningInfo"),
			},
			{
				Name:     "resource_group",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceGroup"),
			},
			{
				Name:     "is_default_container",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsDefaultContainer"),
			},
			{
				Name:     "default_host_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultHostName"),
			},
			{
				Name:     "slot_swap_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SlotSwapStatus"),
			},
			{
				Name:     "https_only",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("HTTPSOnly"),
			},
			{
				Name:     "redundancy_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RedundancyMode"),
			},
			{
				Name:     "storage_account_required",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("StorageAccountRequired"),
			},
			{
				Name:     "key_vault_reference_identity",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyVaultReferenceIdentity"),
			},
			{
				Name:     "virtual_network_subnet_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualNetworkSubnetID"),
			},
			{
				Name:     "identity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Identity"),
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
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},

		Relations: []*schema.Table{
			siteAuthSettings(),
			vnetConnections(),
			publishingProfiles(),
		},
	}
}

func fetchWebApps(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Web.Apps

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
