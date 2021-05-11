package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/compute/v1"
)

func ComputeBackendServices() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_backend_services",
		Resolver:     fetchComputeBackendServices,
		Multiplex:    client.ProjectMultiplex,
		DeleteFilter: client.DeleteProjectFilter,
		IgnoreError:  client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name: "affinity_cookie_ttl_sec",
				Type: schema.TypeBigInt,
			},
			{
				Name:     "cdn_policy_bypass_cache_on_request_headers",
				Type:     schema.TypeStringArray,
				Resolver: resolveComputeBackendServiceCdnPolicyBypassCacheOnRequestHeaders,
			},
			{
				Name:     "cdn_policy_cache_key_policy_include_host",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CdnPolicy.CacheKeyPolicy.IncludeHost"),
			},
			{
				Name:     "cdn_policy_cache_key_policy_include_protocol",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CdnPolicy.CacheKeyPolicy.IncludeProtocol"),
			},
			{
				Name:     "cdn_policy_cache_key_policy_include_query_string",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CdnPolicy.CacheKeyPolicy.IncludeQueryString"),
			},
			{
				Name:     "cdn_policy_cache_key_policy_query_string_blacklist",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CdnPolicy.CacheKeyPolicy.QueryStringBlacklist"),
			},
			{
				Name:     "cdn_policy_cache_key_policy_query_string_whitelist",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CdnPolicy.CacheKeyPolicy.QueryStringWhitelist"),
			},
			{
				Name:     "cdn_policy_cache_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CdnPolicy.CacheMode"),
			},
			{
				Name:     "cdn_policy_client_ttl",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("CdnPolicy.ClientTtl"),
			},
			{
				Name:     "cdn_policy_default_ttl",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("CdnPolicy.DefaultTtl"),
			},
			{
				Name:     "cdn_policy_max_ttl",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("CdnPolicy.MaxTtl"),
			},
			{
				Name:     "cdn_policy_negative_caching",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CdnPolicy.NegativeCaching"),
			},
			{
				Name:     "cdn_policy_request_coalescing",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CdnPolicy.RequestCoalescing"),
			},
			{
				Name:     "cdn_policy_serve_while_stale",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("CdnPolicy.ServeWhileStale"),
			},
			{
				Name:     "cdn_policy_signed_url_cache_max_age_sec",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("CdnPolicy.SignedUrlCacheMaxAgeSec"),
			},
			{
				Name:     "cdn_policy_signed_url_key_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CdnPolicy.SignedUrlKeyNames"),
			},
			{
				Name:     "circuit_breakers_max_connections",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("CircuitBreakers.MaxConnections"),
			},
			{
				Name:     "circuit_breakers_max_pending_requests",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("CircuitBreakers.MaxPendingRequests"),
			},
			{
				Name:     "circuit_breakers_max_requests",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("CircuitBreakers.MaxRequests"),
			},
			{
				Name:     "circuit_breakers_max_requests_per_connection",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("CircuitBreakers.MaxRequestsPerConnection"),
			},
			{
				Name:     "circuit_breakers_max_retries",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("CircuitBreakers.MaxRetries"),
			},
			{
				Name:     "connection_draining_draining_timeout_sec",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ConnectionDraining.DrainingTimeoutSec"),
			},
			{
				Name:     "consistent_hash_http_cookie_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConsistentHash.HttpCookie.Name"),
			},
			{
				Name:     "consistent_hash_http_cookie_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConsistentHash.HttpCookie.Path"),
			},
			{
				Name:     "consistent_hash_http_cookie_ttl_nanos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ConsistentHash.HttpCookie.Ttl.Nanos"),
			},
			{
				Name:     "consistent_hash_http_cookie_ttl_seconds",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ConsistentHash.HttpCookie.Ttl.Seconds"),
			},
			{
				Name:     "consistent_hash_http_header_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConsistentHash.HttpHeaderName"),
			},
			{
				Name:     "consistent_hash_minimum_ring_size",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ConsistentHash.MinimumRingSize"),
			},
			{
				Name: "creation_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "custom_request_headers",
				Type: schema.TypeStringArray,
			},
			{
				Name: "custom_response_headers",
				Type: schema.TypeStringArray,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name:     "enable_cdn",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableCDN"),
			},
			{
				Name:     "failover_policy_disable_connection_drain_on_failover",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("FailoverPolicy.DisableConnectionDrainOnFailover"),
			},
			{
				Name:     "failover_policy_drop_traffic_if_unhealthy",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("FailoverPolicy.DropTrafficIfUnhealthy"),
			},
			{
				Name:     "failover_policy_failover_ratio",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("FailoverPolicy.FailoverRatio"),
			},
			{
				Name: "fingerprint",
				Type: schema.TypeString,
			},
			{
				Name: "health_checks",
				Type: schema.TypeStringArray,
			},
			{
				Name:     "iap_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Iap.Enabled"),
			},
			{
				Name:     "iap_oauth2_client_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Iap.Oauth2ClientId"),
			},
			{
				Name:     "iap_oauth2_client_secret",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Iap.Oauth2ClientSecret"),
			},
			{
				Name:     "iap_oauth2_client_secret_sha256",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Iap.Oauth2ClientSecretSha256"),
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveResourceId,
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name: "load_balancing_scheme",
				Type: schema.TypeString,
			},
			{
				Name: "locality_lb_policy",
				Type: schema.TypeString,
			},
			{
				Name:     "log_config_enable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("LogConfig.Enable"),
			},
			{
				Name:     "log_config_sample_rate",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("LogConfig.SampleRate"),
			},
			{
				Name:     "max_stream_duration_nanos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("MaxStreamDuration.Nanos"),
			},
			{
				Name:     "max_stream_duration_seconds",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("MaxStreamDuration.Seconds"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "network",
				Type: schema.TypeString,
			},
			{
				Name:     "outlier_detection_base_ejection_time_nanos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("OutlierDetection.BaseEjectionTime.Nanos"),
			},
			{
				Name:     "outlier_detection_base_ejection_time_seconds",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("OutlierDetection.BaseEjectionTime.Seconds"),
			},
			{
				Name:     "outlier_detection_consecutive_errors",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("OutlierDetection.ConsecutiveErrors"),
			},
			{
				Name:     "outlier_detection_consecutive_gateway_failure",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("OutlierDetection.ConsecutiveGatewayFailure"),
			},
			{
				Name:     "outlier_detection_enforcing_consecutive_errors",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("OutlierDetection.EnforcingConsecutiveErrors"),
			},
			{
				Name:     "outlier_detection_enforcing_consecutive_gateway_failure",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("OutlierDetection.EnforcingConsecutiveGatewayFailure"),
			},
			{
				Name:     "outlier_detection_enforcing_success_rate",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("OutlierDetection.EnforcingSuccessRate"),
			},
			{
				Name:     "outlier_detection_interval_nanos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("OutlierDetection.Interval.Nanos"),
			},
			{
				Name:     "outlier_detection_interval_seconds",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("OutlierDetection.Interval.Seconds"),
			},
			{
				Name:     "outlier_detection_max_ejection_percent",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("OutlierDetection.MaxEjectionPercent"),
			},
			{
				Name:     "outlier_detection_success_rate_minimum_hosts",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("OutlierDetection.SuccessRateMinimumHosts"),
			},
			{
				Name:     "outlier_detection_success_rate_request_volume",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("OutlierDetection.SuccessRateRequestVolume"),
			},
			{
				Name:     "outlier_detection_success_rate_stdev_factor",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("OutlierDetection.SuccessRateStdevFactor"),
			},
			{
				Name: "port",
				Type: schema.TypeBigInt,
			},
			{
				Name: "port_name",
				Type: schema.TypeString,
			},
			{
				Name: "protocol",
				Type: schema.TypeString,
			},
			{
				Name: "region",
				Type: schema.TypeString,
			},
			{
				Name: "security_policy",
				Type: schema.TypeString,
			},
			{
				Name:     "security_settings_client_tls_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecuritySettings.ClientTlsPolicy"),
			},
			{
				Name:     "security_settings_subject_alt_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SecuritySettings.SubjectAltNames"),
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
			{
				Name: "session_affinity",
				Type: schema.TypeString,
			},
			{
				Name: "timeout_sec",
				Type: schema.TypeBigInt,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_compute_backend_service_backends",
				Resolver: fetchComputeBackendServiceBackends,
				Columns: []schema.Column{
					{
						Name:     "backend_service_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "balancing_mode",
						Type: schema.TypeString,
					},
					{
						Name: "capacity_scaler",
						Type: schema.TypeFloat,
					},
					{
						Name: "description",
						Type: schema.TypeString,
					},
					{
						Name: "failover",
						Type: schema.TypeBool,
					},
					{
						Name: "group",
						Type: schema.TypeString,
					},
					{
						Name: "max_connections",
						Type: schema.TypeBigInt,
					},
					{
						Name: "max_connections_per_endpoint",
						Type: schema.TypeBigInt,
					},
					{
						Name: "max_connections_per_instance",
						Type: schema.TypeBigInt,
					},
					{
						Name: "max_rate",
						Type: schema.TypeBigInt,
					},
					{
						Name: "max_rate_per_endpoint",
						Type: schema.TypeFloat,
					},
					{
						Name: "max_rate_per_instance",
						Type: schema.TypeFloat,
					},
					{
						Name: "max_utilization",
						Type: schema.TypeFloat,
					},
				},
			},
			{
				Name:     "gcp_compute_backend_service_cdn_negative_caching_policies",
				Resolver: fetchComputeBackendServiceCdnNegativeCachingPolicies,
				Columns: []schema.Column{
					{
						Name:     "backend_service_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "code",
						Type: schema.TypeBigInt,
					},
					{
						Name: "ttl",
						Type: schema.TypeBigInt,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeBackendServices(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	nextPageToken := ""
	c := meta.(*client.Client)
	for {
		call := c.Services.Compute.BackendServices.AggregatedList(c.ProjectId).Context(ctx).PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		var backendServices []*compute.BackendService
		for _, backendServicesScopedList := range output.Items {
			backendServices = append(backendServices, backendServicesScopedList.BackendServices...)
		}
		res <- backendServices
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func resolveComputeBackendServiceCdnPolicyBypassCacheOnRequestHeaders(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(*compute.BackendService)
	if r.CdnPolicy == nil {
		return nil
	}
	headers := make([]string, len(r.CdnPolicy.BypassCacheOnRequestHeaders))
	for i, v := range r.CdnPolicy.BypassCacheOnRequestHeaders {
		headers[i] = v.HeaderName
	}
	return resource.Set("cdn_policy_bypass_cache_on_request_headers", headers)
}
func fetchComputeBackendServiceBackends(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(*compute.BackendService)
	res <- r.Backends
	return nil
}
func fetchComputeBackendServiceCdnNegativeCachingPolicies(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(*compute.BackendService)
	if r.CdnPolicy != nil {
		res <- r.CdnPolicy.NegativeCachingPolicy
	}
	return nil
}
