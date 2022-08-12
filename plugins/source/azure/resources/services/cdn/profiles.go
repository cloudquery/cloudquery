package cdn

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource profiles --config profiles.hcl --output .
func Profiles() *schema.Table {
	return &schema.Table{
		Name:         "azure_cdn_profiles",
		Description:  "Profile CDN profile is a logical grouping of endpoints that share the same settings, such as CDN provider and pricing tier",
		Resolver:     fetchCdnProfiles,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription ID",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "sku_name",
				Description: "Name of the pricing tier",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "state",
				Description: "Resource status of the profile",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ProfileProperties.ResourceState"),
			},
			{
				Name:        "provisioning_state",
				Description: "Provisioning status of the profile",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ProfileProperties.ProvisioningState"),
			},
			{
				Name:        "frontdoor_id",
				Description: "The Id of the frontdoor",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ProfileProperties.FrontdoorID"),
			},
			{
				Name:        "location",
				Description: "Resource location",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "id",
				Description: "Resource ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource name",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_by",
				Description: "An identifier for the identity that created the resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SystemData.CreatedBy"),
			},
			{
				Name:        "created_by_type",
				Description: "The type of identity that created the resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SystemData.CreatedByType"),
			},
			{
				Name:     "created_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SystemData.CreatedAt.Time"),
			},
			{
				Name:        "last_modified_by",
				Description: "An identifier for the identity that last modified the resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SystemData.LastModifiedBy"),
			},
			{
				Name:        "last_modified_by_type",
				Description: "The type of identity that last modified the resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SystemData.LastModifiedByType"),
			},
			{
				Name:     "last_modified_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SystemData.LastModifiedAt.Time"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_cdn_profile_endpoints",
				Description: "Endpoint CDN endpoint is the entity within a CDN profile containing configuration information such as origin, protocol, content caching and delivery behavior",
				Resolver:    fetchCdnProfileEndpoints,
				Columns: []schema.Column{
					{
						Name:        "profile_cq_id",
						Description: "Unique CloudQuery ID of azure_cdn_profiles table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "host_name",
						Description: "The host name of the endpoint structured as {endpointName}{DNSZone}, eg",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EndpointProperties.HostName"),
					},
					{
						Name:        "resource_state",
						Description: "Resource status of the endpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EndpointProperties.ResourceState"),
					},
					{
						Name:        "provisioning_state",
						Description: "Provisioning status of the endpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EndpointProperties.ProvisioningState"),
					},
					{
						Name:        "origin_path",
						Description: "A directory path on the origin that CDN can use to retrieve content from, eg",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EndpointProperties.OriginPath"),
					},
					{
						Name:        "content_types_to_compress",
						Description: "List of content types on which compression applies",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("EndpointProperties.ContentTypesToCompress"),
					},
					{
						Name:        "origin_host_header",
						Description: "The host header value sent to the origin with each request",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EndpointProperties.OriginHostHeader"),
					},
					{
						Name:        "is_compression_enabled",
						Description: "Indicates whether content compression is enabled on CDN",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("EndpointProperties.IsCompressionEnabled"),
					},
					{
						Name:        "is_http_allowed",
						Description: "Indicates whether HTTP traffic is allowed on the endpoint",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("EndpointProperties.IsHTTPAllowed"),
					},
					{
						Name:        "is_https_allowed",
						Description: "Indicates whether HTTPS traffic is allowed on the endpoint",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("EndpointProperties.IsHTTPSAllowed"),
					},
					{
						Name:        "query_string_caching_behavior",
						Description: "Defines how CDN caches requests that include query strings",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EndpointProperties.QueryStringCachingBehavior"),
					},
					{
						Name:        "optimization_type",
						Description: "Specifies what scenario the customer wants this CDN endpoint to optimize for, eg",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EndpointProperties.OptimizationType"),
					},
					{
						Name:        "probe_path",
						Description: "Path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EndpointProperties.ProbePath"),
					},
					{
						Name:        "default_origin_group_id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EndpointProperties.DefaultOriginGroup.ID"),
					},
					{
						Name:        "delivery_policy_description",
						Description: "User-friendly description of the policy",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EndpointProperties.DeliveryPolicy.Description"),
					},
					{
						Name:        "web_application_firewall_policy_link_id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EndpointProperties.WebApplicationFirewallPolicyLink.ID"),
					},
					{
						Name:        "location",
						Description: "Resource location",
						Type:        schema.TypeString,
					},
					{
						Name:        "tags",
						Description: "Resource tags",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "Resource name",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Resource type",
						Type:        schema.TypeString,
					},
					{
						Name:        "created_by",
						Description: "An identifier for the identity that created the resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SystemData.CreatedBy"),
					},
					{
						Name:        "created_by_type",
						Description: "The type of identity that created the resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SystemData.CreatedByType"),
					},
					{
						Name:     "created_at_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("SystemData.CreatedAt.Time"),
					},
					{
						Name:        "last_modified_by",
						Description: "An identifier for the identity that last modified the resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SystemData.LastModifiedBy"),
					},
					{
						Name:        "last_modified_by_type",
						Description: "The type of identity that last modified the resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SystemData.LastModifiedByType"),
					},
					{
						Name:     "last_modified_at_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("SystemData.LastModifiedAt.Time"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "azure_cdn_profile_endpoint_origins",
						Description: "DeepCreatedOrigin the main origin of CDN content which is added when creating a CDN endpoint",
						Resolver:    fetchCdnProfileEndpointOrigins,
						Columns: []schema.Column{
							{
								Name:        "profile_endpoint_cq_id",
								Description: "Unique CloudQuery ID of azure_cdn_profile_endpoints table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "Origin name which must be unique within the endpoint",
								Type:        schema.TypeString,
							},
							{
								Name:        "host_name",
								Description: "The address of the origin",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("DeepCreatedOriginProperties.HostName"),
							},
							{
								Name:        "http_port",
								Description: "The value of the HTTP port",
								Type:        schema.TypeBigInt,
								Resolver:    schema.PathResolver("DeepCreatedOriginProperties.HTTPPort"),
							},
							{
								Name:        "https_port",
								Description: "The value of the HTTPS port",
								Type:        schema.TypeBigInt,
								Resolver:    schema.PathResolver("DeepCreatedOriginProperties.HTTPSPort"),
							},
							{
								Name:        "origin_host_header",
								Description: "The host header value sent to the origin with each request",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("DeepCreatedOriginProperties.OriginHostHeader"),
							},
							{
								Name:        "priority",
								Description: "Priority of origin in given origin group for load balancing",
								Type:        schema.TypeBigInt,
								Resolver:    schema.PathResolver("DeepCreatedOriginProperties.Priority"),
							},
							{
								Name:        "weight",
								Description: "Weight of the origin in given origin group for load balancing",
								Type:        schema.TypeBigInt,
								Resolver:    schema.PathResolver("DeepCreatedOriginProperties.Weight"),
							},
							{
								Name:        "enabled",
								Description: "Origin is enabled for load balancing or not",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("DeepCreatedOriginProperties.Enabled"),
							},
							{
								Name:        "private_link_alias",
								Description: "The Alias of the Private Link resource",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("DeepCreatedOriginProperties.PrivateLinkAlias"),
							},
							{
								Name:        "private_link_resource_id",
								Description: "The Resource Id of the Private Link resource",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("DeepCreatedOriginProperties.PrivateLinkResourceID"),
							},
							{
								Name:        "private_link_location",
								Description: "The location of the Private Link resource",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("DeepCreatedOriginProperties.PrivateLinkLocation"),
							},
							{
								Name:        "private_link_approval_message",
								Description: "A custom message to be included in the approval request to connect to the Private Link",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("DeepCreatedOriginProperties.PrivateLinkApprovalMessage"),
							},
						},
					},
					{
						Name:        "azure_cdn_profile_endpoint_origin_groups",
						Description: "DeepCreatedOriginGroup the origin group for CDN content which is added when creating a CDN endpoint Traffic is sent to the origins within the origin group based on origin health",
						Resolver:    fetchCdnProfileEndpointOriginGroups,
						Columns: []schema.Column{
							{
								Name:        "profile_endpoint_cq_id",
								Description: "Unique CloudQuery ID of azure_cdn_profile_endpoints table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "Origin group name which must be unique within the endpoint",
								Type:        schema.TypeString,
							},
							{
								Name:        "health_probe_settings_probe_path",
								Description: "The path relative to the origin that is used to determine the health of the origin",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("DeepCreatedOriginGroupProperties.HealthProbeSettings.ProbePath"),
							},
							{
								Name:        "health_probe_settings_probe_request_type",
								Description: "The type of health probe request that is made",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("DeepCreatedOriginGroupProperties.HealthProbeSettings.ProbeRequestType"),
							},
							{
								Name:        "health_probe_settings_probe_protocol",
								Description: "Protocol to use for health probe",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("DeepCreatedOriginGroupProperties.HealthProbeSettings.ProbeProtocol"),
							},
							{
								Name:        "health_probe_settings_probe_interval_in_seconds",
								Description: "The number of seconds between health probesDefault is 240sec",
								Type:        schema.TypeBigInt,
								Resolver:    schema.PathResolver("DeepCreatedOriginGroupProperties.HealthProbeSettings.ProbeIntervalInSeconds"),
							},
							{
								Name:        "origins",
								Description: "The source of the content being delivered via CDN within given origin group",
								Type:        schema.TypeStringArray,
								Resolver:    resolveProfileEndpointOriginGroupsOrigins,
							},
							{
								Name:        "traffic_restoration_time_to_healed_or_new_endpoints_in_minutes",
								Description: "Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added",
								Type:        schema.TypeBigInt,
								Resolver:    schema.PathResolver("DeepCreatedOriginGroupProperties.TrafficRestorationTimeToHealedOrNewEndpointsInMinutes"),
							},
							{
								Name:        "response_based_origin_error_detection_settings",
								Description: "The JSON object that contains the properties to determine origin health using real requests/responsesThis property is currently not supported",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("DeepCreatedOriginGroupProperties.ResponseBasedOriginErrorDetectionSettings"),
							},
						},
					},
					{
						Name:        "azure_cdn_profile_endpoint_geo_filters",
						Description: "GeoFilter rules defining user's geo access within a CDN endpoint",
						Resolver:    fetchCdnProfileEndpointGeoFilters,
						Columns: []schema.Column{
							{
								Name:        "profile_endpoint_cq_id",
								Description: "Unique CloudQuery ID of azure_cdn_profile_endpoints table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "relative_path",
								Description: "Relative path applicable to geo filter",
								Type:        schema.TypeString,
							},
							{
								Name:        "action",
								Description: "Action of the geo filter, ie",
								Type:        schema.TypeString,
							},
							{
								Name:        "country_codes",
								Description: "Two letter country codes defining user country access in a geo filter, eg",
								Type:        schema.TypeStringArray,
							},
						},
					},
					{
						Name:        "azure_cdn_profile_endpoint_url_signing_keys",
						Description: "URLSigningKey url signing key",
						Resolver:    fetchCdnProfileEndpointUrlSigningKeys,
						Columns: []schema.Column{
							{
								Name:        "profile_endpoint_cq_id",
								Description: "Unique CloudQuery ID of azure_cdn_profile_endpoints table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "key_id",
								Description: "Defines the customer defined key Id",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("KeyID"),
							},
							{
								Name:     "key_source_parameters_odata_type",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("KeySourceParameters.OdataType"),
							},
							{
								Name:        "key_source_parameters_subscription_id",
								Description: "Subscription Id of the user's Key Vault containing the secret",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("KeySourceParameters.SubscriptionID"),
							},
							{
								Name:        "key_source_parameters_resource_group_name",
								Description: "Resource group of the user's Key Vault containing the secret",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("KeySourceParameters.ResourceGroupName"),
							},
							{
								Name:        "key_source_parameters_vault_name",
								Description: "The name of the user's Key Vault containing the secret",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("KeySourceParameters.VaultName"),
							},
							{
								Name:        "key_source_parameters_secret_name",
								Description: "The name of secret in Key Vault",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("KeySourceParameters.SecretName"),
							},
							{
								Name:        "key_source_parameters_secret_version",
								Description: "The version(GUID) of secret in Key Vault",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("KeySourceParameters.SecretVersion"),
							},
						},
					},
					{
						Name:        "azure_cdn_profile_endpoint_delivery_policy_rules",
						Description: "DeliveryRule a rule that specifies a set of actions and conditions",
						Resolver:    fetchCdnProfileEndpointDeliveryPolicyRules,
						Columns: []schema.Column{
							{
								Name:        "profile_endpoint_cq_id",
								Description: "Unique CloudQuery ID of azure_cdn_profile_endpoints table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "Name of the rule",
								Type:        schema.TypeString,
							},
							{
								Name:        "order",
								Description: "The order in which the rules are applied for the endpoint",
								Type:        schema.TypeBigInt,
							},
							{
								Name:        "conditions",
								Description: "A list of conditions that must be matched for the actions to be executed",
								Type:        schema.TypeJSON,
								Resolver:    resolveProfileEndpointDeliveryPolicyRulesConditions,
							},
							{
								Name:        "actions",
								Description: "A list of actions that are executed when all the conditions of a rule are satisfied",
								Type:        schema.TypeJSON,
								Resolver:    resolveProfileEndpointDeliveryPolicyRulesActions,
							},
						},
					},
					{
						Name:        "azure_cdn_profile_endpoint_custom_domains",
						Description: "CustomDomain friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, eg",
						Resolver:    fetchCdnProfileEndpointCustomDomains,
						Columns: []schema.Column{
							{
								Name:        "profile_endpoint_cq_id",
								Description: "Unique CloudQuery ID of azure_cdn_profile_endpoints table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "host_name",
								Description: "The host name of the custom domain",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("CustomDomainProperties.HostName"),
							},
							{
								Name:        "state",
								Description: "Resource status of the custom domain",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("CustomDomainProperties.ResourceState"),
							},
							{
								Name:        "custom_https_provisioning_state",
								Description: "Provisioning status of Custom Https of the custom domain",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("CustomDomainProperties.CustomHTTPSProvisioningState"),
							},
							{
								Name:        "custom_https_provisioning_substate",
								Description: "Provisioning substate shows the progress of custom HTTPS enabling/disabling process step by step",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("CustomDomainProperties.CustomHTTPSProvisioningSubstate"),
							},
							{
								Name:        "custom_https_parameters",
								Description: "Certificate parameters for securing custom HTTPS",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("CustomDomainProperties.CustomHTTPSParameters"),
							},
							{
								Name:        "validation_data",
								Description: "Special validation or data may be required when delivering CDN to some regions due to local compliance reasons",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("CustomDomainProperties.ValidationData"),
							},
							{
								Name:        "provisioning_state",
								Description: "Provisioning status of the custom domain",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("CustomDomainProperties.ProvisioningState"),
							},
							{
								Name:        "id",
								Description: "Resource ID",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ID"),
							},
							{
								Name:        "name",
								Description: "Resource name",
								Type:        schema.TypeString,
							},
							{
								Name:        "type",
								Description: "Resource type",
								Type:        schema.TypeString,
							},
							{
								Name:        "created_by",
								Description: "An identifier for the identity that created the resource",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("SystemData.CreatedBy"),
							},
							{
								Name:        "created_by_type",
								Description: "The type of identity that created the resource",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("SystemData.CreatedByType"),
							},
							{
								Name:     "created_at_time",
								Type:     schema.TypeTimestamp,
								Resolver: schema.PathResolver("SystemData.CreatedAt.Time"),
							},
							{
								Name:        "last_modified_by",
								Description: "An identifier for the identity that last modified the resource",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("SystemData.LastModifiedBy"),
							},
							{
								Name:        "last_modified_by_type",
								Description: "The type of identity that last modified the resource",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("SystemData.LastModifiedByType"),
							},
							{
								Name:     "last_modified_at_time",
								Type:     schema.TypeTimestamp,
								Resolver: schema.PathResolver("SystemData.LastModifiedAt.Time"),
							},
						},
					},
					{
						Name:        "azure_cdn_profile_endpoint_routes",
						Description: "Route friendly Routes name mapping to the any Routes or secret related information",
						Resolver:    fetchCdnProfileEndpointRoutes,
						Columns: []schema.Column{
							{
								Name:        "profile_endpoint_cq_id",
								Description: "Unique CloudQuery ID of azure_cdn_profile_endpoints table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "custom_domains",
								Description: "Domains referenced by this endpoint",
								Type:        schema.TypeStringArray,
								Resolver:    resolveProfileEndpointRoutesCustomDomains,
							},
							{
								Name:        "origin_group_id",
								Description: "Resource ID",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("RouteProperties.OriginGroup.ID"),
							},
							{
								Name:        "origin_path",
								Description: "A directory path on the origin that AzureFrontDoor can use to retrieve content from, eg",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("RouteProperties.OriginPath"),
							},
							{
								Name:        "rule_sets",
								Description: "rule sets referenced by this endpoint",
								Type:        schema.TypeStringArray,
								Resolver:    resolveProfileEndpointRoutesRuleSets,
							},
							{
								Name:        "supported_protocols",
								Description: "List of supported protocols for this route",
								Type:        schema.TypeStringArray,
								Resolver:    schema.PathResolver("RouteProperties.SupportedProtocols"),
							},
							{
								Name:        "patterns_to_match",
								Description: "The route patterns of the rule",
								Type:        schema.TypeStringArray,
								Resolver:    schema.PathResolver("RouteProperties.PatternsToMatch"),
							},
							{
								Name:        "compression_settings",
								Description: "compression settings",
								Type:        schema.TypeJSON,
								Resolver:    resolveProfileEndpointRoutesCompressionSettings,
							},
							{
								Name:        "query_string_caching_behavior",
								Description: "Defines how CDN caches requests that include query strings",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("RouteProperties.QueryStringCachingBehavior"),
							},
							{
								Name:        "forwarding_protocol",
								Description: "Protocol this rule will use when forwarding traffic to backends",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("RouteProperties.ForwardingProtocol"),
							},
							{
								Name:        "link_to_default_domain",
								Description: "whether this route will be linked to the default endpoint domain",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("RouteProperties.LinkToDefaultDomain"),
							},
							{
								Name:        "https_redirect",
								Description: "Whether to automatically redirect HTTP traffic to HTTPS traffic",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("RouteProperties.HTTPSRedirect"),
							},
							{
								Name:        "enabled_state",
								Description: "Whether to enable use of this rule",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("RouteProperties.EnabledState"),
							},
							{
								Name:        "provisioning_state",
								Description: "Provisioning status",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("RouteProperties.ProvisioningState"),
							},
							{
								Name:        "deployment_status",
								Description: "Possible values include: 'DeploymentStatusNotStarted', 'DeploymentStatusInProgress', 'DeploymentStatusSucceeded', 'DeploymentStatusFailed'",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("RouteProperties.DeploymentStatus"),
							},
							{
								Name:        "id",
								Description: "Resource ID",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ID"),
							},
							{
								Name:        "name",
								Description: "Resource name",
								Type:        schema.TypeString,
							},
							{
								Name:        "type",
								Description: "Resource type",
								Type:        schema.TypeString,
							},
							{
								Name:        "created_by",
								Description: "An identifier for the identity that created the resource",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("SystemData.CreatedBy"),
							},
							{
								Name:        "created_by_type",
								Description: "The type of identity that created the resource",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("SystemData.CreatedByType"),
							},
							{
								Name:     "created_at_time",
								Type:     schema.TypeTimestamp,
								Resolver: schema.PathResolver("SystemData.CreatedAt.Time"),
							},
							{
								Name:        "last_modified_by",
								Description: "An identifier for the identity that last modified the resource",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("SystemData.LastModifiedBy"),
							},
							{
								Name:        "last_modified_by_type",
								Description: "The type of identity that last modified the resource",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("SystemData.LastModifiedByType"),
							},
							{
								Name:     "last_modified_at_time",
								Type:     schema.TypeTimestamp,
								Resolver: schema.PathResolver("SystemData.LastModifiedAt.Time"),
							},
						},
					},
				},
			},
			{
				Name:        "azure_cdn_profile_rule_sets",
				Description: "RuleSet friendly RuleSet name mapping to the any RuleSet or secret related information",
				Resolver:    fetchCdnProfileRuleSets,
				Columns: []schema.Column{
					{
						Name:        "profile_cq_id",
						Description: "Unique CloudQuery ID of azure_cdn_profiles table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "provisioning_state",
						Description: "Provisioning status",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RuleSetProperties.ProvisioningState"),
					},
					{
						Name:        "deployment_status",
						Description: "Possible values include: 'DeploymentStatusNotStarted', 'DeploymentStatusInProgress', 'DeploymentStatusSucceeded', 'DeploymentStatusFailed'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RuleSetProperties.DeploymentStatus"),
					},
					{
						Name:        "id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "Resource name",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Resource type",
						Type:        schema.TypeString,
					},
					{
						Name:        "created_by",
						Description: "An identifier for the identity that created the resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SystemData.CreatedBy"),
					},
					{
						Name:        "created_by_type",
						Description: "The type of identity that created the resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SystemData.CreatedByType"),
					},
					{
						Name:     "created_at_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("SystemData.CreatedAt.Time"),
					},
					{
						Name:        "last_modified_by",
						Description: "An identifier for the identity that last modified the resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SystemData.LastModifiedBy"),
					},
					{
						Name:        "last_modified_by_type",
						Description: "The type of identity that last modified the resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SystemData.LastModifiedByType"),
					},
					{
						Name:     "last_modified_at_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("SystemData.LastModifiedAt.Time"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "azure_cdn_profile_rule_set_rules",
						Description: "Rule friendly Rules name mapping to the any Rules or secret related information",
						Resolver:    fetchCdnProfileRuleSetRules,
						Columns: []schema.Column{
							{
								Name:        "profile_rule_set_cq_id",
								Description: "Unique CloudQuery ID of azure_cdn_profile_rule_sets table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "order",
								Description: "The order in which the rules are applied for the endpoint",
								Type:        schema.TypeBigInt,
								Resolver:    schema.PathResolver("RuleProperties.Order"),
							},
							{
								Name:        "conditions",
								Description: "A list of conditions that must be matched for the actions to be executed",
								Type:        schema.TypeJSON,
								Resolver:    resolveProfileRuleSetRulesConditions,
							},
							{
								Name:        "actions",
								Description: "A list of actions that are executed when all the conditions of a rule are satisfied",
								Type:        schema.TypeJSON,
								Resolver:    resolveProfileRuleSetRulesActions,
							},
							{
								Name:        "match_processing_behavior",
								Description: "If this rule is a match should the rules engine continue running the remaining rules or stop",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("RuleProperties.MatchProcessingBehavior"),
							},
							{
								Name:        "provisioning_state",
								Description: "Provisioning status",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("RuleProperties.ProvisioningState"),
							},
							{
								Name:        "deployment_status",
								Description: "Possible values include: 'DeploymentStatusNotStarted', 'DeploymentStatusInProgress', 'DeploymentStatusSucceeded', 'DeploymentStatusFailed'",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("RuleProperties.DeploymentStatus"),
							},
							{
								Name:        "id",
								Description: "Resource ID",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ID"),
							},
							{
								Name:        "name",
								Description: "Resource name",
								Type:        schema.TypeString,
							},
							{
								Name:        "type",
								Description: "Resource type",
								Type:        schema.TypeString,
							},
							{
								Name:        "created_by",
								Description: "An identifier for the identity that created the resource",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("SystemData.CreatedBy"),
							},
							{
								Name:        "created_by_type",
								Description: "The type of identity that created the resource",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("SystemData.CreatedByType"),
							},
							{
								Name:     "created_at_time",
								Type:     schema.TypeTimestamp,
								Resolver: schema.PathResolver("SystemData.CreatedAt.Time"),
							},
							{
								Name:        "last_modified_by",
								Description: "An identifier for the identity that last modified the resource",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("SystemData.LastModifiedBy"),
							},
							{
								Name:        "last_modified_by_type",
								Description: "The type of identity that last modified the resource",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("SystemData.LastModifiedByType"),
							},
							{
								Name:     "last_modified_at_time",
								Type:     schema.TypeTimestamp,
								Resolver: schema.PathResolver("SystemData.LastModifiedAt.Time"),
							},
						},
					},
				},
			},
			{
				Name:        "azure_cdn_profile_security_policies",
				Description: "SecurityPolicy securityPolicy association for AzureFrontDoor profile",
				Resolver:    fetchCdnProfileSecurityPolicies,
				Columns: []schema.Column{
					{
						Name:        "profile_cq_id",
						Description: "Unique CloudQuery ID of azure_cdn_profiles table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "parameters",
						Description: "object which contains security policy parameters",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("SecurityPolicyProperties.Parameters"),
					},
					{
						Name:        "provisioning_state",
						Description: "Provisioning status",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityPolicyProperties.ProvisioningState"),
					},
					{
						Name:        "deployment_status",
						Description: "Possible values include: 'DeploymentStatusNotStarted', 'DeploymentStatusInProgress', 'DeploymentStatusSucceeded', 'DeploymentStatusFailed'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityPolicyProperties.DeploymentStatus"),
					},
					{
						Name:        "id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "Resource name",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Resource type",
						Type:        schema.TypeString,
					},
					{
						Name:        "created_by",
						Description: "An identifier for the identity that created the resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SystemData.CreatedBy"),
					},
					{
						Name:        "created_by_type",
						Description: "The type of identity that created the resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SystemData.CreatedByType"),
					},
					{
						Name:     "created_at_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("SystemData.CreatedAt.Time"),
					},
					{
						Name:        "last_modified_by",
						Description: "An identifier for the identity that last modified the resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SystemData.LastModifiedBy"),
					},
					{
						Name:        "last_modified_by_type",
						Description: "The type of identity that last modified the resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SystemData.LastModifiedByType"),
					},
					{
						Name:     "last_modified_at_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("SystemData.LastModifiedAt.Time"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCdnProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().CDN.Profiles
	response, err := svc.List(ctx)
	if err != nil {
		return diag.WrapError(err)
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
func fetchCdnProfileEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(cdn.Profile)
	svc := meta.(*client.Client).Services().CDN.Endpoints
	details, err := client.ParseResourceID(*r.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	response, err := svc.ListByProfile(ctx, details.ResourceGroup, details.ResourceName)
	if err != nil {
		return diag.WrapError(err)
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
func fetchCdnProfileEndpointOrigins(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(cdn.Endpoint)
	if r.Origins == nil {
		return nil
	}
	res <- *r.Origins
	return nil
}
func fetchCdnProfileEndpointOriginGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(cdn.Endpoint)
	if r.OriginGroups == nil {
		return nil
	}
	res <- *r.OriginGroups
	return nil
}
func resolveProfileEndpointOriginGroupsOrigins(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(cdn.DeepCreatedOriginGroup)
	if r.Origins == nil {
		return nil
	}
	sa := make([]string, 0, len(*r.Origins))
	for _, i := range *r.Origins {
		sa = append(sa, *i.ID)
	}
	return diag.WrapError(resource.Set(c.Name, sa))
}
func fetchCdnProfileEndpointGeoFilters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(cdn.Endpoint)
	if r.GeoFilters == nil {
		return nil
	}
	res <- *r.GeoFilters
	return nil
}
func fetchCdnProfileEndpointUrlSigningKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(cdn.Endpoint)
	if r.URLSigningKeys == nil {
		return nil
	}
	res <- *r.URLSigningKeys
	return nil
}
func fetchCdnProfileEndpointDeliveryPolicyRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(cdn.Endpoint)
	if r.DeliveryPolicy == nil || r.DeliveryPolicy.Rules == nil {
		return nil
	}
	res <- *r.DeliveryPolicy.Rules
	return nil
}
func resolveProfileEndpointDeliveryPolicyRulesConditions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(cdn.DeliveryRule)
	if r.Conditions == nil {
		return nil
	}
	j, err := marshalConditions(*r.Conditions)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func resolveProfileEndpointDeliveryPolicyRulesActions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(cdn.DeliveryRule)
	if r.Actions == nil {
		return nil
	}
	j, err := marshalActions(*r.Actions)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func fetchCdnProfileEndpointCustomDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Parent.Item.(cdn.Profile)
	e := parent.Item.(cdn.Endpoint)
	svc := meta.(*client.Client).Services().CDN.CustomDomains
	details, err := client.ParseResourceID(*r.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	response, err := svc.ListByEndpoint(ctx, details.ResourceGroup, details.ResourceName, *e.Name)
	if err != nil {
		return diag.WrapError(err)
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
func fetchCdnProfileEndpointRoutes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Parent.Item.(cdn.Profile)
	e := parent.Item.(cdn.Endpoint)
	svc := meta.(*client.Client).Services().CDN.Routes
	details, err := client.ParseResourceID(*r.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	response, err := svc.ListByEndpoint(ctx, details.ResourceGroup, details.ResourceName, *e.Name)
	if err != nil {
		return diag.WrapError(err)
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
func resolveProfileEndpointRoutesCustomDomains(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(cdn.Route)
	if r.CustomDomains == nil {
		return nil
	}
	sa := make([]string, 0, len(*r.CustomDomains))
	for _, i := range *r.CustomDomains {
		sa = append(sa, *i.ID)
	}
	return diag.WrapError(resource.Set(c.Name, sa))
}
func resolveProfileEndpointRoutesRuleSets(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(cdn.Route)
	if r.RuleSets == nil {
		return nil
	}
	sa := make([]string, 0, len(*r.RuleSets))
	for _, i := range *r.RuleSets {
		sa = append(sa, *i.ID)
	}
	return diag.WrapError(resource.Set(c.Name, sa))
}
func resolveProfileEndpointRoutesCompressionSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(cdn.Route)
	if r.CompressionSettings == nil {
		return nil
	}
	b, err := json.Marshal(r.CompressionSettings)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}
func fetchCdnProfileRuleSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(cdn.Profile)
	svc := meta.(*client.Client).Services().CDN.RuleSets
	details, err := client.ParseResourceID(*r.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	response, err := svc.ListByProfile(ctx, details.ResourceGroup, details.ResourceName)
	if err != nil {
		return diag.WrapError(err)
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
func fetchCdnProfileRuleSetRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Parent.Item.(cdn.Profile)
	r := parent.Item.(cdn.RuleSet)
	svc := meta.(*client.Client).Services().CDN.Rules
	details, err := client.ParseResourceID(*p.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	response, err := svc.ListByRuleSet(ctx, details.ResourceGroup, details.ResourceName, *r.Name)
	if err != nil {
		return diag.WrapError(err)
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
func resolveProfileRuleSetRulesConditions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(cdn.Rule)
	if r.Conditions == nil {
		return nil
	}
	j, err := marshalConditions(*r.Conditions)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func resolveProfileRuleSetRulesActions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(cdn.Rule)
	if r.Actions == nil {
		return nil
	}
	j, err := marshalActions(*r.Actions)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func fetchCdnProfileSecurityPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(cdn.Profile)
	svc := meta.(*client.Client).Services().CDN.SecurityPolicies
	details, err := client.ParseResourceID(*r.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	response, err := svc.ListByProfile(ctx, details.ResourceGroup, details.ResourceName)
	if err != nil {
		return diag.WrapError(err)
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func resolveCdnProfileSecurityPolicyWebApplicationFirewallParameters(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(cdn.SecurityPolicy)
	waf, ok := r.Parameters.AsSecurityPolicyParameters()
	if !ok {
		return nil
	}
	return diag.WrapError(resource.Set(c.Name, waf))
}
func resolveCdnProfileSecurityPolicyParametersType(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(cdn.SecurityPolicy)
	params, ok := r.Parameters.AsSecurityPolicyParameters()
	if !ok {
		return nil
	}

	return diag.WrapError(resource.Set(c.Name, params))
}
