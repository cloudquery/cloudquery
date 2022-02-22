package iothub

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/iothub/mgmt/2021-07-02/devices"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IothubHubs() *schema.Table {
	return &schema.Table{
		Name:         "azure_iothub_hubs",
		Description:  "Azure IoT hub.",
		Resolver:     fetchIothubHubs,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "etag",
				Description: "The Etag.",
				Type:        schema.TypeString,
			},
			{
				Name:        "disable_local_auth",
				Description: "If true, SAS tokens with Iot hub scoped SAS keys cannot be used for authentication.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Properties.DisableLocalAuth"),
			},
			{
				Name:        "disable_device_sas",
				Description: "If true, all device(including Edge devices but excluding modules) scoped SAS keys cannot be used for authentication.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Properties.DisableDeviceSAS"),
			},
			{
				Name:        "disable_module_sas",
				Description: "If true, all module scoped SAS keys cannot be used for authentication.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Properties.DisableModuleSAS"),
			},
			{
				Name:        "restrict_outbound_network_access",
				Description: "If true, egress from IotHub will be restricted to only the allowed FQDNs that are configured via allowedFqdnList.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Properties.RestrictOutboundNetworkAccess"),
			},
			{
				Name:        "allowed_fqdn_list",
				Description: "List of allowed FQDNs(Fully Qualified Domain Name) for egress from Iot Hub.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Properties.AllowedFqdnList"),
			},
			{
				Name:        "public_network_access",
				Description: "Whether requests from Public Network are allowed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.PublicNetworkAccess"),
			},
			{
				Name:        "network_rule_sets_default_action",
				Description: "Default Action for Network Rule Set",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.NetworkRuleSets.DefaultAction"),
			},
			{
				Name:        "network_rule_sets_apply_to_built_in_event_hub_endpoint",
				Description: "If True, then Network Rule Set is also applied to BuiltIn EventHub EndPoint of IotHub",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Properties.NetworkRuleSets.ApplyToBuiltInEventHubEndpoint"),
			},
			{
				Name:        "min_tls_version",
				Description: "Specifies the minimum TLS version to support for this hub",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.MinTLSVersion"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.ProvisioningState"),
			},
			{
				Name:        "state",
				Description: "The hub state.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.State"),
			},
			{
				Name:        "host_name",
				Description: "The name of the host.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.HostName"),
			},
			{
				Name:        "event_hub_endpoints",
				Description: "The Event Hub-compatible endpoint properties",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Properties.EventHubEndpoints"),
			},
			{
				Name:        "routing_fallback_route_name",
				Description: "The name of the route",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.Routing.FallbackRoute.Name"),
			},
			{
				Name:        "routing_fallback_route_source",
				Description: "The source to which the routing rule is to be applied to",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.Routing.FallbackRoute.Source"),
			},
			{
				Name:        "routing_fallback_route_condition",
				Description: "The condition which is evaluated in order to apply the fallback route",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.Routing.FallbackRoute.Condition"),
			},
			{
				Name:        "routing_fallback_route_endpoint_names",
				Description: "The list of endpoints to which the messages that satisfy the condition are routed to",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Properties.Routing.FallbackRoute.EndpointNames"),
			},
			{
				Name:        "routing_fallback_route_is_enabled",
				Description: "Used to specify whether the fallback route is enabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Properties.Routing.FallbackRoute.IsEnabled"),
			},
			{
				Name:        "routing_enrichments",
				Description: "The list of user-provided enrichments that the IoT hub applies to messages to be delivered to built-in and custom endpoints",
				Type:        schema.TypeJSON,
				Resolver:    resolveIothubHubsRoutingEnrichments,
			},
			{
				Name:        "storage_endpoints",
				Description: "The list of Azure Storage endpoints where you can upload files",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Properties.StorageEndpoints"),
			},
			{
				Name:        "messaging_endpoints",
				Description: "The messaging endpoint properties for the file upload notification queue.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Properties.MessagingEndpoints"),
			},
			{
				Name:        "enable_file_upload_notifications",
				Description: "If True, file upload notifications are enabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Properties.EnableFileUploadNotifications"),
			},
			{
				Name:        "cloud_to_device_max_delivery_count",
				Description: "The max delivery count for cloud-to-device messages in the device queue",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Properties.CloudToDevice.MaxDeliveryCount"),
			},
			{
				Name:        "cloud_to_device_default_ttl_as_iso8601",
				Description: "The default time to live for cloud-to-device messages in the device queue",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.CloudToDevice.DefaultTTLAsIso8601"),
			},
			{
				Name:        "cloud_to_device_feedback_lock_duration_as_iso8601",
				Description: "The lock duration for the feedback queue",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.CloudToDevice.Feedback.LockDurationAsIso8601"),
			},
			{
				Name:        "cloud_to_device_feedback_ttl_as_iso8601",
				Description: "The period of time for which a message is available to consume before it is expired by the IoT hub",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.CloudToDevice.Feedback.TTLAsIso8601"),
			},
			{
				Name:        "cloud_to_device_feedback_max_delivery_count",
				Description: "The number of times the IoT hub attempts to deliver a message on the feedback queue",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Properties.CloudToDevice.Feedback.MaxDeliveryCount"),
			},
			{
				Name:        "comments",
				Description: "IoT hub comments.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.Comments"),
			},
			{
				Name:        "features",
				Description: "The capabilities and features enabled for the IoT hub",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.Features"),
			},
			{
				Name:        "locations",
				Description: "Primary and secondary location for iot hub",
				Type:        schema.TypeJSON,
				Resolver:    resolveIothubHubsLocations,
			},
			{
				Name:        "enable_data_residency",
				Description: "This property when set to true, will enable data residency, thus, disabling disaster recovery.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Properties.EnableDataResidency"),
			},
			{
				Name:        "sku_name",
				Description: "The name of the SKU",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "The billing tier for the IoT hub",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Tier"),
			},
			{
				Name:        "sku_capacity",
				Description: "The number of provisioned IoT Hub units",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Sku.Capacity"),
			},
			{
				Name:        "identity_principal_id",
				Description: "Principal Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.PrincipalID"),
			},
			{
				Name:        "identity_tenant_id",
				Description: "Tenant Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.TenantID"),
			},
			{
				Name:        "identity_type",
				Description: "The type of identity used for the resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
			{
				Name:        "identity_user_assigned_identities",
				Type:        schema.TypeJSON,
				Description: "The identities of assigned users",
				Resolver:    schema.PathResolver("Identity.UserAssignedIdentities"),
			},
			{
				Name:        "system_data_created_by",
				Description: "The identity that created the resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SystemData.CreatedBy"),
			},
			{
				Name:        "system_data_created_by_type",
				Description: "The type of identity that created the resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SystemData.CreatedByType"),
			},
			{
				Name:        "system_data_created_at_time",
				Type:        schema.TypeTimestamp,
				Description: "Created time",
				Resolver:    schema.PathResolver("SystemData.CreatedAt.Time"),
			},
			{
				Name:        "system_data_last_modified_by",
				Description: "The identity that last modified the resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SystemData.LastModifiedBy"),
			},
			{
				Name:        "system_data_last_modified_by_type",
				Description: "The type of identity that last modified the resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SystemData.LastModifiedByType"),
			},
			{
				Name:        "system_data_last_modified_at_time",
				Type:        schema.TypeTimestamp,
				Description: "Last modified time",
				Resolver:    schema.PathResolver("SystemData.LastModifiedAt.Time"),
			},
			{
				Name:        "id",
				Description: "The resource identifier.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "The resource location.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The resource tags.",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_iothub_hub_authorization_policies",
				Description: "SharedAccessSignatureAuthorizationRule the properties of an IoT hub shared access policy.",
				Resolver:    fetchIothubHubAuthorizationPolicies,
				Columns: []schema.Column{
					{
						Name:        "hub_cq_id",
						Description: "Unique CloudQuery ID of azure_iothub_hubs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "key_name",
						Description: "The name of the shared access policy.",
						Type:        schema.TypeString,
					},
					{
						Name:        "primary_key",
						Description: "The primary key.",
						Type:        schema.TypeString,
					},
					{
						Name:        "secondary_key",
						Description: "The secondary key.",
						Type:        schema.TypeString,
					},
					{
						Name:        "rights",
						Description: "The permissions assigned to the shared access policy",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_iothub_hub_ip_filter_rules",
				Description: "IPFilterRule the IP filter rules for the IoT hub.",
				Resolver:    fetchIothubHubIpFilterRules,
				Columns: []schema.Column{
					{
						Name:        "hub_cq_id",
						Description: "Unique CloudQuery ID of azure_iothub_hubs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "filter_name",
						Description: "The name of the IP filter rule.",
						Type:        schema.TypeString,
					},
					{
						Name:        "action",
						Description: "The desired action for requests captured by this rule",
						Type:        schema.TypeString,
					},
					{
						Name:        "ip_mask",
						Description: "A string that contains the IP address range in CIDR notation for the rule.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IPMask"),
					},
				},
			},
			{
				Name:        "azure_iothub_hub_network_rule_sets_ip_rules",
				Description: "NetworkRuleSetIPRule IP Rule to be applied as part of Network Rule Set",
				Resolver:    fetchIothubHubNetworkRuleSetsIpRules,
				Columns: []schema.Column{
					{
						Name:        "hub_cq_id",
						Description: "Unique CloudQuery ID of azure_iothub_hubs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "filter_name",
						Description: "Name of the IP filter rule.",
						Type:        schema.TypeString,
					},
					{
						Name:        "action",
						Description: "IP Filter Action",
						Type:        schema.TypeString,
					},
					{
						Name:        "ip_mask",
						Description: "A string that contains the IP address range in CIDR notation for the rule.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IPMask"),
					},
				},
			},
			{
				Name:        "azure_iothub_hub_private_endpoint_connections",
				Description: "PrivateEndpointConnection the private endpoint connection of an IotHub",
				Resolver:    fetchIothubHubPrivateEndpointConnections,
				Columns: []schema.Column{
					{
						Name:        "hub_cq_id",
						Description: "Unique CloudQuery ID of azure_iothub_hubs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "The resource identifier.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "The resource name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The resource type.",
						Type:        schema.TypeString,
					},
					{
						Name:        "private_endpoint_id",
						Description: "The resource identifier.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateEndpoint.ID"),
					},
					{
						Name:        "status",
						Description: "The status of a private endpoint connection",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateLinkServiceConnectionState.Status"),
					},
					{
						Name:        "description",
						Description: "The description for the current state of a private endpoint connection",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateLinkServiceConnectionState.Description"),
					},
					{
						Name:        "actions_required",
						Description: "Actions required for a private endpoint connection",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateLinkServiceConnectionState.ActionsRequired"),
					},
				},
			},
			{
				Name:        "azure_iothub_hub_routing_endpoints_service_bus_queues",
				Description: "RoutingServiceBusQueueEndpointProperties the properties related to service bus queue endpoint types.",
				Resolver:    fetchIothubHubRoutingEndpointsServiceBusQueues,
				Columns: []schema.Column{
					{
						Name:        "hub_cq_id",
						Description: "Unique CloudQuery ID of azure_iothub_hubs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "Id of the service bus queue endpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "connection_string",
						Description: "The connection string of the service bus queue endpoint.",
						Type:        schema.TypeString,
					},
					{
						Name:        "endpoint_uri",
						Description: "The url of the service bus queue endpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EndpointURI"),
					},
					{
						Name:        "entity_path",
						Description: "Queue name on the service bus namespace",
						Type:        schema.TypeString,
					},
					{
						Name:        "authentication_type",
						Description: "Method used to authenticate against the service bus queue endpoint",
						Type:        schema.TypeString,
					},
					{
						Name:        "identity_user_assigned_identity",
						Description: "The user assigned identity.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Identity.UserAssignedIdentity"),
					},
					{
						Name:        "name",
						Description: "The name that identifies this endpoint",
						Type:        schema.TypeString,
					},
					{
						Name:        "subscription_id",
						Description: "The subscription identifier of the service bus queue endpoint.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubscriptionID"),
					},
					{
						Name:        "resource_group",
						Description: "The name of the resource group of the service bus queue endpoint.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_iothub_hub_routing_endpoints_service_bus_topics",
				Description: "RoutingServiceBusTopicEndpointProperties the properties related to service bus topic endpoint types.",
				Resolver:    fetchIothubHubRoutingEndpointsServiceBusTopics,
				Columns: []schema.Column{
					{
						Name:        "hub_cq_id",
						Description: "Unique CloudQuery ID of azure_iothub_hubs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "Id of the service bus topic endpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "connection_string",
						Description: "The connection string of the service bus topic endpoint.",
						Type:        schema.TypeString,
					},
					{
						Name:        "endpoint_uri",
						Description: "The url of the service bus topic endpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EndpointURI"),
					},
					{
						Name:        "entity_path",
						Description: "Queue name on the service bus topic",
						Type:        schema.TypeString,
					},
					{
						Name:        "authentication_type",
						Description: "Method used to authenticate against the service bus topic endpoint",
						Type:        schema.TypeString,
					},
					{
						Name:        "identity_user_assigned_identity",
						Description: "The user assigned identity.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Identity.UserAssignedIdentity"),
					},
					{
						Name:        "name",
						Description: "The name that identifies this endpoint",
						Type:        schema.TypeString,
					},
					{
						Name:        "subscription_id",
						Description: "The subscription identifier of the service bus topic endpoint.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubscriptionID"),
					},
					{
						Name:        "resource_group",
						Description: "The name of the resource group of the service bus topic endpoint.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_iothub_hub_routing_endpoints_event_hubs",
				Description: "RoutingEventHubProperties the properties related to an event hub endpoint.",
				Resolver:    fetchIothubHubRoutingEndpointsEventHubs,
				Columns: []schema.Column{
					{
						Name:        "hub_cq_id",
						Description: "Unique CloudQuery ID of azure_iothub_hubs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "Id of the event hub endpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "connection_string",
						Description: "The connection string of the event hub endpoint.",
						Type:        schema.TypeString,
					},
					{
						Name:        "endpoint_uri",
						Description: "The url of the event hub endpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EndpointURI"),
					},
					{
						Name:        "entity_path",
						Description: "Event hub name on the event hub namespace",
						Type:        schema.TypeString,
					},
					{
						Name:        "authentication_type",
						Description: "Method used to authenticate against the event hub endpoint",
						Type:        schema.TypeString,
					},
					{
						Name:        "identity_user_assigned_identity",
						Description: "The user assigned identity.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Identity.UserAssignedIdentity"),
					},
					{
						Name:        "name",
						Description: "The name that identifies this endpoint",
						Type:        schema.TypeString,
					},
					{
						Name:        "subscription_id",
						Description: "The subscription identifier of the event hub endpoint.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubscriptionID"),
					},
					{
						Name:        "resource_group",
						Description: "The name of the resource group of the event hub endpoint.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_iothub_hub_routing_endpoints_storage_containers",
				Description: "RoutingStorageContainerProperties the properties related to a storage container endpoint.",
				Resolver:    fetchIothubHubRoutingEndpointsStorageContainers,
				Columns: []schema.Column{
					{
						Name:        "hub_cq_id",
						Description: "Unique CloudQuery ID of azure_iothub_hubs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "Id of the storage container endpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "connection_string",
						Description: "The connection string of the storage account.",
						Type:        schema.TypeString,
					},
					{
						Name:        "endpoint_uri",
						Description: "The url of the storage endpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EndpointURI"),
					},
					{
						Name:        "authentication_type",
						Description: "Method used to authenticate against the storage endpoint",
						Type:        schema.TypeString,
					},
					{
						Name:        "identity_user_assigned_identity",
						Description: "The user assigned identity.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Identity.UserAssignedIdentity"),
					},
					{
						Name:        "name",
						Description: "The name that identifies this endpoint",
						Type:        schema.TypeString,
					},
					{
						Name:        "subscription_id",
						Description: "The subscription identifier of the storage account.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubscriptionID"),
					},
					{
						Name:        "resource_group",
						Description: "The name of the resource group of the storage account.",
						Type:        schema.TypeString,
					},
					{
						Name:        "container_name",
						Description: "The name of storage container in the storage account.",
						Type:        schema.TypeString,
					},
					{
						Name:        "file_name_format",
						Description: "File name format for the blob",
						Type:        schema.TypeString,
					},
					{
						Name:        "batch_frequency_in_seconds",
						Description: "Time interval at which blobs are written to storage",
						Type:        schema.TypeInt,
					},
					{
						Name:        "max_chunk_size_in_bytes",
						Description: "Maximum number of bytes for each blob written to storage",
						Type:        schema.TypeInt,
					},
					{
						Name:        "encoding",
						Description: "Encoding that is used to serialize messages to blobs",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_iothub_hub_routing_routes",
				Description: "RouteProperties the properties of a routing rule that your IoT hub uses to route messages to endpoints.",
				Resolver:    fetchIothubHubRoutingRoutes,
				Columns: []schema.Column{
					{
						Name:        "hub_cq_id",
						Description: "Unique CloudQuery ID of azure_iothub_hubs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "The name of the route",
						Type:        schema.TypeString,
					},
					{
						Name:        "source",
						Description: "The source that the routing rule is to be applied to, such as DeviceMessages",
						Type:        schema.TypeString,
					},
					{
						Name:        "condition",
						Description: "The condition that is evaluated to apply the routing rule",
						Type:        schema.TypeString,
					},
					{
						Name:        "endpoint_names",
						Description: "The list of endpoints to which messages that satisfy the condition are routed",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "is_enabled",
						Description: "Used to specify whether a route is enabled.",
						Type:        schema.TypeBool,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchIothubHubs(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().IotHub
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
func resolveIothubHubsRoutingEnrichments(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	iothub, ok := resource.Item.(devices.IotHubDescription)
	if !ok {
		return fmt.Errorf("expected to have devices.IotHubDescription but got %T", resource.Item)
	}
	if iothub.Properties == nil || iothub.Properties.Routing == nil || iothub.Properties.Routing.Enrichments == nil {
		return nil
	}
	b, err := json.Marshal(iothub.Properties.Routing.Enrichments)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveIothubHubsLocations(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	iothub, ok := resource.Item.(devices.IotHubDescription)
	if !ok {
		return fmt.Errorf("expected to have devices.IotHubDescription but got %T", resource.Item)
	}
	if iothub.Properties == nil || iothub.Properties.Locations == nil {
		return nil
	}
	b, err := json.Marshal(iothub.Properties.Locations)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func fetchIothubHubAuthorizationPolicies(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	iothub, ok := parent.Item.(devices.IotHubDescription)
	if !ok {
		return fmt.Errorf("expected to have devices.IotHubDescription but got %T", parent.Item)
	}
	if iothub.Properties == nil || iothub.Properties.AuthorizationPolicies == nil {
		return nil
	}
	res <- *iothub.Properties.AuthorizationPolicies
	return nil
}
func fetchIothubHubIpFilterRules(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	iothub, ok := parent.Item.(devices.IotHubDescription)
	if !ok {
		return fmt.Errorf("expected to have devices.IotHubDescription but got %T", parent.Item)
	}
	if iothub.Properties == nil || iothub.Properties.IPFilterRules == nil {
		return nil
	}
	res <- *iothub.Properties.IPFilterRules
	return nil
}
func fetchIothubHubNetworkRuleSetsIpRules(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	iothub, ok := parent.Item.(devices.IotHubDescription)
	if !ok {
		return fmt.Errorf("expected to have devices.IotHubDescription but got %T", parent.Item)
	}
	if iothub.Properties == nil || iothub.Properties.NetworkRuleSets == nil || iothub.Properties.NetworkRuleSets.IPRules == nil {
		return nil
	}
	res <- *iothub.Properties.NetworkRuleSets.IPRules
	return nil
}
func fetchIothubHubPrivateEndpointConnections(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	iothub, ok := parent.Item.(devices.IotHubDescription)
	if !ok {
		return fmt.Errorf("expected to have devices.IotHubDescription but got %T", parent.Item)
	}
	if iothub.Properties == nil || iothub.Properties.PrivateEndpointConnections == nil {
		return nil
	}
	res <- *iothub.Properties.PrivateEndpointConnections
	return nil
}
func fetchIothubHubRoutingEndpointsServiceBusQueues(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	iothub, ok := parent.Item.(devices.IotHubDescription)
	if !ok {
		return fmt.Errorf("expected to have devices.IotHubDescription but got %T", parent.Item)
	}
	if iothub.Properties == nil || iothub.Properties.Routing == nil || iothub.Properties.Routing.Endpoints == nil || iothub.Properties.Routing.Endpoints.ServiceBusQueues == nil {
		return nil
	}
	res <- *iothub.Properties.Routing.Endpoints.ServiceBusQueues
	return nil
}
func fetchIothubHubRoutingEndpointsServiceBusTopics(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	iothub, ok := parent.Item.(devices.IotHubDescription)
	if !ok {
		return fmt.Errorf("expected to have devices.IotHubDescription but got %T", parent.Item)
	}
	if iothub.Properties == nil || iothub.Properties.Routing == nil || iothub.Properties.Routing.Endpoints == nil || iothub.Properties.Routing.Endpoints.ServiceBusTopics == nil {
		return nil
	}
	res <- *iothub.Properties.Routing.Endpoints.ServiceBusTopics
	return nil
}
func fetchIothubHubRoutingEndpointsEventHubs(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	iothub, ok := parent.Item.(devices.IotHubDescription)
	if !ok {
		return fmt.Errorf("expected to have devices.IotHubDescription but got %T", parent.Item)
	}
	if iothub.Properties == nil || iothub.Properties.Routing == nil || iothub.Properties.Routing.Endpoints == nil || iothub.Properties.Routing.Endpoints.EventHubs == nil {
		return nil
	}
	res <- *iothub.Properties.Routing.Endpoints.EventHubs
	return nil
}
func fetchIothubHubRoutingEndpointsStorageContainers(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	iothub, ok := parent.Item.(devices.IotHubDescription)
	if !ok {
		return fmt.Errorf("expected to have devices.IotHubDescription but got %T", parent.Item)
	}
	if iothub.Properties == nil || iothub.Properties.Routing == nil || iothub.Properties.Routing.Endpoints == nil || iothub.Properties.Routing.Endpoints.StorageContainers == nil {
		return nil
	}
	res <- *iothub.Properties.Routing.Endpoints.StorageContainers
	return nil
}
func fetchIothubHubRoutingRoutes(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	iothub, ok := parent.Item.(devices.IotHubDescription)
	if !ok {
		return fmt.Errorf("expected to have devices.IotHubDescription but got %T", parent.Item)
	}
	if iothub.Properties == nil || iothub.Properties.Routing == nil || iothub.Properties.Routing.Routes == nil {
		return nil
	}
	res <- *iothub.Properties.Routing.Routes
	return nil
}
