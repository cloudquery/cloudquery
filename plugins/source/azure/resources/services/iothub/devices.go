// Auto generated code - DO NOT EDIT.

package iothub

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Devices() *schema.Table {
	return &schema.Table{
		Name:      "azure_iothub_devices",
		Resolver:  fetchIotHubDevices,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "authorization_policies",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.AuthorizationPolicies"),
			},
			{
				Name:     "disable_local_auth",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.DisableLocalAuth"),
			},
			{
				Name:     "disable_device_sas",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.DisableDeviceSAS"),
			},
			{
				Name:     "disable_module_sas",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.DisableModuleSAS"),
			},
			{
				Name:     "restrict_outbound_network_access",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.RestrictOutboundNetworkAccess"),
			},
			{
				Name:     "allowed_fqdn_list",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Properties.AllowedFqdnList"),
			},
			{
				Name:     "public_network_access",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.PublicNetworkAccess"),
			},
			{
				Name:     "ip_filter_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.IPFilterRules"),
			},
			{
				Name:     "network_rule_sets",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.NetworkRuleSets"),
			},
			{
				Name:     "min_tls_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.MinTLSVersion"),
			},
			{
				Name:     "private_endpoint_connections",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.PrivateEndpointConnections"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.ProvisioningState"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.State"),
			},
			{
				Name:     "host_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.HostName"),
			},
			{
				Name:     "event_hub_endpoints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.EventHubEndpoints"),
			},
			{
				Name:     "routing",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.Routing"),
			},
			{
				Name:     "storage_endpoints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.StorageEndpoints"),
			},
			{
				Name:     "messaging_endpoints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.MessagingEndpoints"),
			},
			{
				Name:     "enable_file_upload_notifications",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.EnableFileUploadNotifications"),
			},
			{
				Name:     "cloud_to_device",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.CloudToDevice"),
			},
			{
				Name:     "comments",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.Comments"),
			},
			{
				Name:     "features",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.Features"),
			},
			{
				Name:     "locations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.Locations"),
			},
			{
				Name:     "enable_data_residency",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.EnableDataResidency"),
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
			},
			{
				Name:     "identity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Identity"),
			},
			{
				Name:     "system_data",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SystemData"),
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

func fetchIotHubDevices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().IotHub.Devices

	response, err := svc.ListBySubscription(ctx)

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
