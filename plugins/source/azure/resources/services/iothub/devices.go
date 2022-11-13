// Auto generated code - DO NOT EDIT.

package iothub

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Devices() *schema.Table {
	return &schema.Table{
		Name:        "azure_iothub_devices",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/iothub/mgmt/2021-07-02/devices#IotHubDescription`,
		Resolver:    fetchIotHubDevices,
		Multiplex:   client.SubscriptionMultiplex,
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
				Name:     "properties_authorization_policies",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.AuthorizationPolicies"),
			},
			{
				Name:     "properties_disable_local_auth",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.DisableLocalAuth"),
			},
			{
				Name:     "properties_disable_device_sas",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.DisableDeviceSAS"),
			},
			{
				Name:     "properties_disable_module_sas",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.DisableModuleSAS"),
			},
			{
				Name:     "properties_restrict_outbound_network_access",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.RestrictOutboundNetworkAccess"),
			},
			{
				Name:     "properties_allowed_fqdn_list",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Properties.AllowedFqdnList"),
			},
			{
				Name:     "properties_public_network_access",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.PublicNetworkAccess"),
			},
			{
				Name:     "properties_ip_filter_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.IPFilterRules"),
			},
			{
				Name:     "properties_network_rule_sets",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.NetworkRuleSets"),
			},
			{
				Name:     "properties_min_tls_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.MinTLSVersion"),
			},
			{
				Name:     "properties_private_endpoint_connections",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.PrivateEndpointConnections"),
			},
			{
				Name:     "properties_provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.ProvisioningState"),
			},
			{
				Name:     "properties_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.State"),
			},
			{
				Name:     "properties_host_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.HostName"),
			},
			{
				Name:     "properties_event_hub_endpoints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.EventHubEndpoints"),
			},
			{
				Name:     "properties_routing",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.Routing"),
			},
			{
				Name:     "properties_storage_endpoints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.StorageEndpoints"),
			},
			{
				Name:     "properties_messaging_endpoints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.MessagingEndpoints"),
			},
			{
				Name:     "properties_enable_file_upload_notifications",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.EnableFileUploadNotifications"),
			},
			{
				Name:     "properties_cloud_to_device",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.CloudToDevice"),
			},
			{
				Name:     "properties_comments",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.Comments"),
			},
			{
				Name:     "properties_features",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.Features"),
			},
			{
				Name:     "properties_locations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.Locations"),
			},
			{
				Name:     "properties_enable_data_residency",
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
