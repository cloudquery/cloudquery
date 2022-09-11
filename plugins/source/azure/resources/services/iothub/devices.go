// Auto generated code - DO NOT EDIT.

package iothub

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
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
				Resolver: schema.PathResolver("AuthorizationPolicies"),
			},
			{
				Name:     "disable_local_auth",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DisableLocalAuth"),
			},
			{
				Name:     "disable_device_sas",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DisableDeviceSAS"),
			},
			{
				Name:     "disable_module_sas",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DisableModuleSAS"),
			},
			{
				Name:     "restrict_outbound_network_access",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("RestrictOutboundNetworkAccess"),
			},
			{
				Name:     "allowed_fqdn_list",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AllowedFqdnList"),
			},
			{
				Name:     "public_network_access",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicNetworkAccess"),
			},
			{
				Name:     "ip_filter_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IPFilterRules"),
			},
			{
				Name:     "network_rule_sets",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkRuleSets"),
			},
			{
				Name:     "min_tls_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MinTLSVersion"),
			},
			{
				Name:     "private_endpoint_connections",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrivateEndpointConnections"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "host_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HostName"),
			},
			{
				Name:     "event_hub_endpoints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EventHubEndpoints"),
			},
			{
				Name:     "routing",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Routing"),
			},
			{
				Name:     "storage_endpoints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StorageEndpoints"),
			},
			{
				Name:     "messaging_endpoints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MessagingEndpoints"),
			},
			{
				Name:     "enable_file_upload_notifications",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableFileUploadNotifications"),
			},
			{
				Name:     "cloud_to_device",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CloudToDevice"),
			},
			{
				Name:     "comments",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Comments"),
			},
			{
				Name:     "features",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Features"),
			},
			{
				Name:     "locations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Locations"),
			},
			{
				Name:     "enable_data_residency",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableDataResidency"),
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
