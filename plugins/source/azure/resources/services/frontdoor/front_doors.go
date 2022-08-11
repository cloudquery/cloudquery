package frontdoor

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-11-01/frontdoor"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource front_doors --config gen.hcl --output .
func FrontDoors() *schema.Table {
	return &schema.Table{
		Name:         "azure_front_doors",
		Description:  "Front Door represents a collection of backend endpoints to route traffic to along with rules that specify how traffic is sent there.",
		Resolver:     fetchFrontDoors,
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
				Name:        "resource_state",
				Description: "Resource state of the Front Door",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.ResourceState"),
			},
			{
				Name:        "provisioning_state",
				Description: "Provisioning state of the Front Door",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.ProvisioningState"),
			},
			{
				Name:        "cname",
				Description: "The host that each frontend endpoint must CNAME to",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.Cname"),
			},
			{
				Name:        "frontdoor_id",
				Description: "The ID of the Front Door",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.FrontdoorID"),
			},
			{
				Name:        "friendly_name",
				Description: "A friendly name for the Front Door",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.FriendlyName"),
			},
			{
				Name:        "enforce_certificate_name_check",
				Description: "Whether to enforce certificate name check on HTTPS requests to all backend pools",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.BackendPoolsSettings.EnforceCertificateNameCheck"),
			},
			{
				Name:        "send_recv_timeout_seconds",
				Description: "Send and receive timeout on forwarding request to the backend",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Properties.BackendPoolsSettings.SendRecvTimeoutSeconds"),
			},
			{
				Name:        "enabled_state",
				Description: "Operational status of the Front Door load balancer",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.EnabledState"),
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
				Name:        "location",
				Description: "Resource location",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "azure_front_door_rules_engines",
				Description:   "Rules engine configuration containing a list of rules that will run to modify the runtime behavior of the request and response.",
				Resolver:      fetchFrontDoorFrontDoorRulesEngines,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "front_door_cq_id",
						Description: "Unique CloudQuery ID of azure_front_doors table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "resource_state",
						Description: "Resource state",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RulesEngineProperties.ResourceState"),
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
						Name:        "id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "azure_front_door_rules_engine_rules",
						Description: "A list of rules that define a particular Rules Engine Configuration.",
						Resolver:    fetchFrontDoorFrontDoorRulesEngineRules,
						Columns: []schema.Column{
							{
								Name:        "front_door_rules_engine_cq_id",
								Description: "Unique CloudQuery ID of azure_front_door_rules_engines table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "A name to refer to the rule",
								Type:        schema.TypeString,
							},
							{
								Name:        "priority",
								Description: "A priority assigned to the rule",
								Type:        schema.TypeInt,
							},
							{
								Name:        "request_header_actions",
								Description: "A list of header actions to apply from the request from AFD to the origin.",
								Type:        schema.TypeJSON,
								Resolver:    resolveFrontDoorRulesEngineRulesRequestHeaderActions,
							},
							{
								Name:        "response_header_actions",
								Description: "A list of header actions to apply from the response from AFD to the client.",
								Type:        schema.TypeJSON,
								Resolver:    resolveFrontDoorRulesEngineRulesResponseHeaderActions,
							},
							{
								Name:        "route_configuration_override",
								Description: "Override the route configuration",
								Type:        schema.TypeJSON,
								Resolver:    resolveFrontDoorRulesEngineRulesRouteConfigurationOverride,
							},
							{
								Name:        "match_processing_behavior",
								Description: "If the rule is a match should the rules engine continue running the remaining rules or stop",
								Type:        schema.TypeString,
							},
						},
						Relations: []*schema.Table{
							{
								Name:        "azure_front_door_rules_engine_rule_match_conditions",
								Description: "A list of match conditions that must meet in order for the actions of the rule to run. Having no match conditions means the actions will always run.",
								Resolver:    fetchFrontDoorFrontDoorRulesEngineRuleMatchConditions,
								Columns: []schema.Column{
									{
										Name:        "front_door_rules_engine_rule_cq_id",
										Description: "Unique CloudQuery ID of azure_front_door_rules_engine_rules table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "match_variable",
										Description: "Match variable",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("RulesEngineMatchVariable"),
									},
									{
										Name:        "selector",
										Description: "Name of selector in request header or request body to be matched",
										Type:        schema.TypeString,
									},
									{
										Name:        "operator",
										Description: "Describes operator to apply to the match condition",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("RulesEngineOperator"),
									},
									{
										Name:        "negate_condition",
										Description: "Describes if this is negate condition or not",
										Type:        schema.TypeBool,
									},
									{
										Name:        "match_value",
										Description: "Match values to match against",
										Type:        schema.TypeStringArray,
										Resolver:    schema.PathResolver("RulesEngineMatchValue"),
									},
									{
										Name:        "transforms",
										Description: "List of transforms",
										Type:        schema.TypeStringArray,
									},
								},
							},
						},
					},
				},
			},
			{
				Name:        "azure_front_door_routing_rules",
				Description: "Routing rules represent specifications for traffic to treat and where to send it, along with health probe information.",
				Resolver:    fetchFrontdoorFrontDoorRoutingRules,
				Columns: []schema.Column{
					{
						Name:        "front_door_cq_id",
						Description: "Unique CloudQuery ID of azure_front_doors table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "resource_state",
						Description: "Resource state",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RoutingRuleProperties.ResourceState"),
					},
					{
						Name:        "frontend_endpoints",
						Description: "Frontend endpoints associated with the rule",
						Type:        schema.TypeStringArray,
						Resolver:    resolveFrontDoorRoutingRulesFrontendEndpoints,
					},
					{
						Name:        "accepted_protocols",
						Description: "Protocol schemes to match for the rule",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("RoutingRuleProperties.AcceptedProtocols"),
					},
					{
						Name:        "patterns_to_match",
						Description: "The route patterns of the rule",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("RoutingRuleProperties.PatternsToMatch"),
					},
					{
						Name:        "enabled_state",
						Description: "Whether the rule is enabled",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RoutingRuleProperties.EnabledState"),
					},
					{
						Name:        "route_configuration",
						Description: "A reference to the routing configuration",
						Type:        schema.TypeJSON,
						Resolver:    resolveFrontDoorRoutingRulesRouteConfiguration,
					},
					{
						Name:          "rules_engine_id",
						Description:   "ID of a specific Rules Engine Configuration to apply to the route",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("RoutingRuleProperties.RulesEngine.ID"),
						IgnoreInTests: true,
					},
					{
						Name:          "web_application_firewall_policy_link_id",
						Description:   "ID of the Web Application Firewall policy for each routing rule (if applicable)",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("RoutingRuleProperties.WebApplicationFirewallPolicyLink.ID"),
						IgnoreInTests: true,
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
						Name:        "id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
			},
			{
				Name:        "azure_front_door_load_balancing_settings",
				Description: "Load balancing settings for a backend pool associated with the Front Door instance",
				Resolver:    fetchFrontdoorFrontDoorLoadBalancingSettings,
				Columns: []schema.Column{
					{
						Name:        "front_door_cq_id",
						Description: "Unique CloudQuery ID of azure_front_doors table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "resource_state",
						Description: "Resource state",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LoadBalancingSettingsProperties.ResourceState"),
					},
					{
						Name:        "sample_size",
						Description: "The number of samples to consider for load balancing decisions",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LoadBalancingSettingsProperties.SampleSize"),
					},
					{
						Name:        "successful_samples_required",
						Description: "The number of samples within the sample period that must succeed",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LoadBalancingSettingsProperties.SuccessfulSamplesRequired"),
					},
					{
						Name:        "additional_latency_milliseconds",
						Description: "The additional latency in milliseconds for probes to fall into the lowest latency bucket",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LoadBalancingSettingsProperties.AdditionalLatencyMilliseconds"),
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
						Name:        "id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
			},
			{
				Name:        "azure_front_door_health_probe_settings",
				Description: "Health probe settings for a backend pool associated with this Front Door instance",
				Resolver:    fetchFrontdoorFrontDoorHealthProbeSettings,
				Columns: []schema.Column{
					{
						Name:        "front_door_cq_id",
						Description: "Unique CloudQuery ID of azure_front_doors table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "resource_state",
						Description: "Resource state",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("HealthProbeSettingsProperties.ResourceState"),
					},
					{
						Name:        "path",
						Description: "The path to use for the health probe",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("HealthProbeSettingsProperties.Path"),
					},
					{
						Name:        "protocol",
						Description: "Protocol scheme to use for the health probe",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("HealthProbeSettingsProperties.Protocol"),
					},
					{
						Name:        "interval_in_seconds",
						Description: "The number of seconds between health probes",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("HealthProbeSettingsProperties.IntervalInSeconds"),
					},
					{
						Name:        "health_probe_method",
						Description: "Which HTTP method is used to perform the health probe",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("HealthProbeSettingsProperties.HealthProbeMethod"),
					},
					{
						Name:        "enabled_state",
						Description: "Whether the health probe is enabled",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("HealthProbeSettingsProperties.EnabledState"),
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
						Name:        "id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
			},
			{
				Name:        "azure_front_door_backend_pools",
				Description: "Backend pools available to routing rules",
				Resolver:    fetchFrontdoorFrontDoorBackendPools,
				Columns: []schema.Column{
					{
						Name:        "front_door_cq_id",
						Description: "Unique CloudQuery ID of azure_front_doors table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "resource_state",
						Description: "Resource state",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("BackendPoolProperties.ResourceState"),
					},
					{
						Name:        "load_balancing_settings_id",
						Description: "Load balancing settings ID for the backend pool",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("BackendPoolProperties.LoadBalancingSettings.ID"),
					},
					{
						Name:        "health_probe_settings_id",
						Description: "L7 health probe settings ID for the backend pool",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("BackendPoolProperties.HealthProbeSettings.ID"),
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
						Name:        "id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "azure_front_door_backend_pool_backends",
						Description: "The set of backends for the backend pool",
						Resolver:    fetchFrontdoorFrontDoorBackendPoolBackends,
						Columns: []schema.Column{
							{
								Name:        "front_door_backend_pool_cq_id",
								Description: "Unique CloudQuery ID of azure_front_door_backend_pools table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "address",
								Description: "Location of the backend (IP address or FQDN)",
								Type:        schema.TypeString,
							},
							{
								Name:          "private_link_alias",
								Description:   "The Alias of the Private Link resource",
								Type:          schema.TypeString,
								IgnoreInTests: true,
							},
							{
								Name:          "private_link_resource_id",
								Description:   "The Resource ID of the Private Link resource",
								Type:          schema.TypeString,
								Resolver:      schema.PathResolver("PrivateLinkResourceID"),
								IgnoreInTests: true,
							},
							{
								Name:          "private_link_location",
								Description:   "The location of the Private Link resource",
								Type:          schema.TypeString,
								IgnoreInTests: true,
							},
							{
								Name:        "private_endpoint_status",
								Description: "The Approval status for the connection to the Private Link",
								Type:        schema.TypeString,
							},
							{
								Name:          "private_link_approval_message",
								Description:   "A custom message to be included in the approval request to connect to the Private Link",
								Type:          schema.TypeString,
								IgnoreInTests: true,
							},
							{
								Name:        "http_port",
								Description: "The HTTP TCP port number",
								Type:        schema.TypeInt,
								Resolver:    schema.PathResolver("HTTPPort"),
							},
							{
								Name:        "https_port",
								Description: "The HTTPS TCP port number",
								Type:        schema.TypeInt,
								Resolver:    schema.PathResolver("HTTPSPort"),
							},
							{
								Name:        "enabled_state",
								Description: "Whether the use of the backend is enabled",
								Type:        schema.TypeString,
							},
							{
								Name:        "priority",
								Description: "Priority to use for load balancing",
								Type:        schema.TypeInt,
							},
							{
								Name:        "weight",
								Description: "Weight of the endpoint for load balancing purposes",
								Type:        schema.TypeInt,
							},
							{
								Name:        "host_header",
								Description: "The value to use as the host header sent to the backend",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("BackendHostHeader"),
							},
						},
					},
				},
			},
			{
				Name:        "azure_front_door_frontend_endpoints",
				Description: "Frontend endpoints available to routing rules",
				Resolver:    fetchFrontdoorFrontDoorFrontendEndpoints,
				Columns: []schema.Column{
					{
						Name:        "front_door_cq_id",
						Description: "Unique CloudQuery ID of azure_front_doors table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "resource_state",
						Description: "Resource state",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.ResourceState"),
					},
					{
						Name:        "custom_https_provisioning_state",
						Description: "Provisioning status of custom https of the frontend endpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.CustomHTTPSProvisioningState"),
					},
					{
						Name:        "custom_https_provisioning_substate",
						Description: "Provisioning substate shows the progress of custom HTTPS enabling/disabling process step by step",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.CustomHTTPSProvisioningSubstate"),
					},
					{
						Name:        "certificate_source",
						Description: "Defines the source of the SSL certificate",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.CustomHTTPSConfiguration.CertificateSource"),
					},
					{
						Name:          "protocol_type",
						Description:   "Defines the TLS extension protocol that is used for secure delivery",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("FrontendEndpointProperties.CustomHTTPSConfiguration.ProtocolType"),
						IgnoreInTests: true,
					},
					{
						Name:        "minimum_tls_version",
						Description: "The minimum TLS version required from the clients to establish an SSL handshake with Front Door",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.CustomHTTPSConfiguration.MinimumTLSVersion"),
					},
					{
						Name:          "vault_id",
						Description:   "The Key Vault containing the SSL certificate",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("FrontendEndpointProperties.CustomHTTPSConfiguration.KeyVaultCertificateSourceParameters.Vault.ID"),
						IgnoreInTests: true,
					},
					{
						Name:          "secret_name",
						Description:   "The name of the Key Vault secret representing the full certificate PFX",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("FrontendEndpointProperties.CustomHTTPSConfiguration.KeyVaultCertificateSourceParameters.SecretName"),
						IgnoreInTests: true,
					},
					{
						Name:          "secret_version",
						Description:   "The version of the Key Vault secret representing the full certificate PFX",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("FrontendEndpointProperties.CustomHTTPSConfiguration.KeyVaultCertificateSourceParameters.SecretVersion"),
						IgnoreInTests: true,
					},
					{
						Name:        "certificate_type",
						Description: "The type of the certificate used for secure connections to the frontend endpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.CustomHTTPSConfiguration.CertificateSourceParameters.CertificateType"),
					},
					{
						Name:        "host_name",
						Description: "The host name of the frontend endpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.HostName"),
					},
					{
						Name:        "session_affinity_enabled_state",
						Description: "Whether session affinity is allowed on the host",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.SessionAffinityEnabledState"),
					},
					{
						Name:        "session_affinity_ttl_seconds",
						Description: "The TTL to use in seconds for session affinity, if applicable",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.SessionAffinityTTLSeconds"),
					},
					{
						Name:          "web_application_firewall_policy_link_id",
						Description:   "Defines the Web Application Firewall policy for each host (if applicable)",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("FrontendEndpointProperties.WebApplicationFirewallPolicyLink.ID"),
						IgnoreInTests: true,
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
						Name:        "id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchFrontDoors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().FrontDoor
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
func fetchFrontDoorFrontDoorRulesEngines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	frontDoor := parent.Item.(frontdoor.FrontDoor)
	if frontDoor.RulesEngines != nil {
		res <- *frontDoor.RulesEngines
	}
	return nil
}
func fetchFrontDoorFrontDoorRulesEngineRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	rulesEngine := parent.Item.(frontdoor.RulesEngine)
	if rulesEngine.Rules != nil {
		res <- *rulesEngine.Rules
	}
	return nil
}
func resolveFrontDoorRulesEngineRulesRequestHeaderActions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule := resource.Item.(frontdoor.RulesEngineRule)
	if rule.Action == nil || rule.Action.RequestHeaderActions == nil {
		return nil
	}

	data, err := marshalHeaderActions(*rule.Action.RequestHeaderActions)
	if err != nil {
		return err
	}
	return diag.WrapError(resource.Set("request_header_actions", data))
}
func resolveFrontDoorRulesEngineRulesResponseHeaderActions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule := resource.Item.(frontdoor.RulesEngineRule)
	if rule.Action == nil || rule.Action.ResponseHeaderActions == nil {
		return nil
	}

	data, err := marshalHeaderActions(*rule.Action.ResponseHeaderActions)
	if err != nil {
		return err
	}

	return diag.WrapError(resource.Set("response_header_actions", data))
}
func resolveFrontDoorRulesEngineRulesRouteConfigurationOverride(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule := resource.Item.(frontdoor.RulesEngineRule)
	if rule.Action == nil || rule.Action.RouteConfigurationOverride == nil {
		return nil
	}

	data, err := marshalRouteConfiguration(rule.Action.RouteConfigurationOverride)
	if err != nil {
		return err
	}

	return diag.WrapError(resource.Set("route_configuration_override", data))
}
func fetchFrontDoorFrontDoorRulesEngineRuleMatchConditions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	rule := parent.Item.(frontdoor.RulesEngineRule)
	if rule.MatchConditions != nil {
		res <- *rule.MatchConditions
	}
	return nil
}
func fetchFrontdoorFrontDoorRoutingRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	frontDoor := parent.Item.(frontdoor.FrontDoor)
	if frontDoor.RoutingRules != nil {
		res <- *frontDoor.RoutingRules
	}
	return nil
}
func resolveFrontDoorRoutingRulesFrontendEndpoints(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule := resource.Item.(frontdoor.RoutingRule)
	if rule.FrontendEndpoints == nil {
		return nil
	}
	var endpoints []string
	for _, ep := range *rule.FrontendEndpoints {
		if ep.ID != nil {
			endpoints = append(endpoints, *ep.ID)
		}
	}
	return diag.WrapError(resource.Set("frontend_endpoints", endpoints))
}
func resolveFrontDoorRoutingRulesRouteConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule := resource.Item.(frontdoor.RoutingRule)
	if rule.RouteConfiguration == nil {
		return nil
	}

	data, err := marshalRouteConfiguration(rule.RouteConfiguration)
	if err != nil {
		return err
	}

	return diag.WrapError(resource.Set("route_configuration", data))
}
func fetchFrontdoorFrontDoorLoadBalancingSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	frontDoor := parent.Item.(frontdoor.FrontDoor)
	if frontDoor.LoadBalancingSettings != nil {
		res <- *frontDoor.LoadBalancingSettings
	}
	return nil
}
func fetchFrontdoorFrontDoorHealthProbeSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	frontDoor := parent.Item.(frontdoor.FrontDoor)
	if frontDoor.HealthProbeSettings != nil {
		res <- *frontDoor.HealthProbeSettings
	}
	return nil
}
func fetchFrontdoorFrontDoorBackendPools(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	frontDoor := parent.Item.(frontdoor.FrontDoor)
	if frontDoor.BackendPools != nil {
		res <- *frontDoor.BackendPools
	}
	return nil
}
func fetchFrontdoorFrontDoorBackendPoolBackends(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	backendPool := parent.Item.(frontdoor.BackendPool)
	if backendPool.Backends != nil {
		res <- *backendPool.Backends
	}
	return nil
}
func fetchFrontdoorFrontDoorFrontendEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	frontDoor := parent.Item.(frontdoor.FrontDoor)
	if frontDoor.FrontendEndpoints != nil {
		res <- *frontDoor.FrontendEndpoints
	}
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func marshalHeaderActions(actions []frontdoor.HeaderAction) (data []byte, err error) {
	type headerActionJSON struct {
		ActionType frontdoor.HeaderActionType `json:"action_type"`
		HeaderName *string                    `json:"header_name,omitempty"`
		Value      *string                    `json:"value,omitempty"`
	}

	raw := make([]json.RawMessage, 0, len(actions))
	for _, action := range actions {
		actionJSON := headerActionJSON{
			ActionType: action.HeaderActionType,
			HeaderName: action.HeaderName,
			Value:      action.Value,
		}
		data, err = json.Marshal(actionJSON)
		if err != nil {
			return nil, diag.WrapError(err)
		}
		raw = append(raw, data)
	}

	data, err = json.Marshal(raw)
	return data, diag.WrapError(err)
}
func marshalRouteConfiguration(config frontdoor.BasicRouteConfiguration) (data []byte, err error) {
	dataMessage := map[string]json.RawMessage{}
	if route, ok := config.AsRouteConfiguration(); ok && route != nil {
		data, err := route.MarshalJSON()
		if err != nil {
			return nil, diag.WrapError(err)
		}
		dataMessage["route_configuration"] = data
	}

	if forward, ok := config.AsForwardingConfiguration(); ok && forward != nil {
		data, err := forward.MarshalJSON()
		if err != nil {
			return nil, diag.WrapError(err)
		}
		dataMessage["forwarding_configuration"] = data
	}

	if redirect, ok := config.AsRedirectConfiguration(); ok && redirect != nil {
		data, err := redirect.MarshalJSON()
		if err != nil {
			return nil, diag.WrapError(err)
		}
		dataMessage["redirect_configuration"] = data
	}

	data, err = json.Marshal(dataMessage)
	return data, diag.WrapError(err)
}
