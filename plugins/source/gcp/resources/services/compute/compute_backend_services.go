package compute

import (
	"context"
	"encoding/json"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/compute/v1"
)

func ComputeBackendServices() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_backend_services",
		Description: "Represents a Backend Service resource  A backend service defines how Google Cloud load balancers distribute traffic The backend service configuration contains a set of values, such as the protocol used to connect to backends, various distribution and session settings, health checks, and timeouts These settings provide fine-grained control over how your load balancer behaves.",
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "id"}},

		Resolver:  fetchComputeBackendServices,
		Multiplex: client.ProjectMultiplex,

		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "affinity_cookie_ttl_sec",
				Description: "Lifetime of cookies in seconds Only applicable if the loadBalancingScheme is EXTERNAL, INTERNAL_SELF_MANAGED, or INTERNAL_MANAGED, the protocol is HTTP or HTTPS, and the sessionAffinity is GENERATED_COOKIE, or HTTP_COOKIE  If set to 0, the cookie is non-persistent and lasts only until the end of the browser session (or equivalent) The maximum allowed value is one day (86,400)  Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true",
				Type:        schema.TypeBigInt,
			},
			{
				Name:          "cdn_policy_bypass_cache_on_request_headers",
				Description:   "Bypass the cache when the specified request headers are matched - eg Pragma or Authorization headers Up to 5 headers can be specified The cache is bypassed for all cdnPolicycacheMode settings",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
				Resolver:      resolveComputeBackendServiceCdnPolicyBypassCacheOnRequestHeaders,
			},
			{
				Name:        "cdn_policy_cache_key_policy_include_host",
				Description: "If true, requests to different hosts will be cached separately",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("CdnPolicy.CacheKeyPolicy.IncludeHost"),
			},
			{
				Name:        "cdn_policy_cache_key_policy_include_protocol",
				Description: "If true, http and https requests will be cached separately",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("CdnPolicy.CacheKeyPolicy.IncludeProtocol"),
			},
			{
				Name:        "cdn_policy_cache_key_policy_include_query_string",
				Description: "If true, include query string parameters in the cache key according to query_string_whitelist and query_string_blacklist If neither is set, the entire query string will be included If false, the query string will be excluded from the cache key entirely",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("CdnPolicy.CacheKeyPolicy.IncludeQueryString"),
			},
			{
				Name:          "cdn_policy_cache_key_policy_query_string_blacklist",
				Description:   "Names of query string parameters to exclude in cache keys All other parameters will be included Either specify query_string_whitelist or query_string_blacklist, not both '&' and '=' will be percent encoded and not treated as delimiters",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("CdnPolicy.CacheKeyPolicy.QueryStringBlacklist"),
			},
			{
				Name:          "cdn_policy_cache_key_policy_query_string_whitelist",
				Description:   "Names of query string parameters to include in cache keys All other parameters will be excluded Either specify query_string_whitelist or query_string_blacklist, not both '&' and '=' will be percent encoded and not treated as delimiters",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("CdnPolicy.CacheKeyPolicy.QueryStringWhitelist"),
			},
			{
				Name:        "cdn_policy_cache_mode",
				Description: "Specifies the cache setting for all responses from this backend",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CdnPolicy.CacheMode"),
			},
			{
				Name:        "cdn_policy_client_ttl",
				Description: "Specifies a separate client (eg browser client) maximum TTL",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("CdnPolicy.ClientTtl"),
			},
			{
				Name:        "cdn_policy_default_ttl",
				Description: "Specifies the default TTL for cached content served by this origin for responses that do not have an existing valid TTL",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("CdnPolicy.DefaultTtl"),
			},
			{
				Name:        "cdn_policy_max_ttl",
				Description: "Specifies the maximum allowed TTL for cached content served by this origin Cache directives that attempt to set a max-age or s-maxage higher than this, or an Expires header more than maxTTL seconds in the future will be capped at the value of maxTTL, as if it were the value of an s-maxage Cache-Control directive Headers sent to the client will not be modified Setting a TTL of \"0\" means \"always revalidate\" The maximum allowed value is 31,622,400s (1 year), noting that infrequently accessed objects may be evicted from the cache before the defined TTL",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("CdnPolicy.MaxTtl"),
			},
			{
				Name:        "cdn_policy_negative_caching",
				Description: "Negative caching allows per-status code TTLs to be set, in order to apply fine-grained caching for common errors or redirects This can reduce the load on your origin and improve end-user experience by reducing response latency When the cache mode is set to CACHE_ALL_STATIC or USE_ORIGIN_HEADERS, negative caching applies to responses with the specified response code that lack any Cache-Control, Expires, or Pragma: no-cache directives When the cache mode is set to FORCE_CACHE_ALL, negative caching applies to all responses with the specified response code, and override any caching headers By default, Cloud CDN will apply the following default TTLs to these status codes: HTTP 300 (Multiple Choice), 301, 308 (Permanent Redirects): 10m HTTP 404 (Not Found), 410 (Gone), 451 (Unavailable For Legal Reasons): 120s HTTP 405 (Method Not Found), 421 (Misdirected Request), 501 (Not Implemented): 60s These defaults can be overridden in negative_caching_policy",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("CdnPolicy.NegativeCaching"),
			},
			{
				Name:          "cdn_policy_negative_caching_policy",
				Description:   "Sets a cache TTL for the specified HTTP status code negative_caching must be enabled to configure negative_caching_policy Omitting the policy and leaving negative_caching enabled will use Cloud CDN's default cache TTLs Note that when specifying an explicit negative_caching_policy, you should take care to specify a cache TTL for all response codes that you wish to cache Cloud CDN will not apply any default negative caching when a policy exists",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
				Resolver:      resolveComputeBackendServiceCdnPolicyNegativeCachingPolicy,
			},
			{
				Name:        "cdn_policy_request_coalescing",
				Description: "If true then Cloud CDN will combine multiple concurrent cache fill requests into a small number of requests to the origin",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("CdnPolicy.RequestCoalescing"),
			},
			{
				Name:        "cdn_policy_serve_while_stale",
				Description: "Serve existing content from the cache (if available) when revalidating content with the origin, or when an error is encountered when refreshing the cache This setting defines the default \"max-stale\" duration for any cached responses that do not specify a max-stale directive Stale responses that exceed the TTL configured here will not be served The default limit (max-stale) is 86400s (1 day), which will allow stale content to be served up to this limit beyond the max-age (or s-max-age) of a cached response The maximum allowed value is 604800 (1 week) Set this to zero (0) to disable serve-while-stale",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("CdnPolicy.ServeWhileStale"),
			},
			{
				Name:        "cdn_policy_signed_url_cache_max_age_sec",
				Description: "Maximum number of seconds the response to a signed URL request will be considered fresh After this time period, the response will be revalidated before being served Defaults to 1hr (3600s) When serving responses to signed URL requests, Cloud CDN will internally behave as though all responses from this backend had a \"Cache-Control: public, max-age=[TTL]\" header, regardless of any existing Cache-Control header The actual headers served in responses will not be altered",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("CdnPolicy.SignedUrlCacheMaxAgeSec"),
			},
			{
				Name:          "cdn_policy_signed_url_key_names",
				Description:   "Names of the keys for signing request URLs",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("CdnPolicy.SignedUrlKeyNames"),
			},
			{
				Name:        "circuit_breakers_max_connections",
				Description: "The maximum number of connections to the backend service If not specified, there is no limit",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("CircuitBreakers.MaxConnections"),
			},
			{
				Name:        "circuit_breakers_max_pending_requests",
				Description: "The maximum number of pending requests allowed to the backend service If not specified, there is no limit",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("CircuitBreakers.MaxPendingRequests"),
			},
			{
				Name:        "circuit_breakers_max_requests",
				Description: "The maximum number of parallel requests that allowed to the backend service If not specified, there is no limit",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("CircuitBreakers.MaxRequests"),
			},
			{
				Name:        "circuit_breakers_max_requests_per_connection",
				Description: "Maximum requests for a single connection to the backend service This parameter is respected by both the HTTP/11 and HTTP/2 implementations If not specified, there is no limit Setting this parameter to 1 will effectively disable keep alive",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("CircuitBreakers.MaxRequestsPerConnection"),
			},
			{
				Name:        "circuit_breakers_max_retries",
				Description: "The maximum number of parallel retries allowed to the backend cluster If not specified, the default is 1",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("CircuitBreakers.MaxRetries"),
			},
			{
				Name:        "connection_draining_draining_timeout_sec",
				Description: "Configures a duration timeout for existing requests on a removed backend instance For supported load balancers and protocols, as described in Enabling connection draining",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ConnectionDraining.DrainingTimeoutSec"),
			},
			{
				Name:        "consistent_hash_http_cookie_name",
				Description: "Name of the cookie",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ConsistentHash.HttpCookie.Name"),
			},
			{
				Name:        "consistent_hash_http_cookie_path",
				Description: "Path to set for the cookie",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ConsistentHash.HttpCookie.Path"),
			},
			{
				Name:        "consistent_hash_http_cookie_ttl_nanos",
				Description: "Span of time that's a fraction of a second at nanosecond resolution Durations less than one second are represented with a 0 `seconds` field and a positive `nanos` field Must be from 0 to 999,999,999 inclusive",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ConsistentHash.HttpCookie.Ttl.Nanos"),
			},
			{
				Name:        "consistent_hash_http_cookie_ttl_seconds",
				Description: "Span of time at a resolution of a second Must be from 0 to 315,576,000,000 inclusive Note: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ConsistentHash.HttpCookie.Ttl.Seconds"),
			},
			{
				Name:        "consistent_hash_http_header_name",
				Description: "The hash based on the value of the specified header field This field is applicable if the sessionAffinity is set to HEADER_FIELD",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ConsistentHash.HttpHeaderName"),
			},
			{
				Name:        "consistent_hash_minimum_ring_size",
				Description: "The minimum number of virtual nodes to use for the hash ring Defaults to 1024 Larger ring sizes result in more granular load distributions If the number of hosts in the load balancing pool is larger than the ring size, each host will be assigned a single virtual node",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ConsistentHash.MinimumRingSize"),
			},
			{
				Name:        "creation_timestamp",
				Description: "Creation timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:          "custom_request_headers",
				Description:   "Headers that the HTTP/S load balancer should add to proxied requests",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:          "custom_response_headers",
				Description:   "Headers that the HTTP/S load balancer should add to proxied responses",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "description",
				Description: "An optional description of this resource Provide this property when you create the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "enable_cdn",
				Description: "If true, enables Cloud CDN for the backend service Only applicable if the loadBalancingScheme is EXTERNAL and the protocol is HTTP or HTTPS",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("EnableCDN"),
			},
			{
				Name:        "failover_policy_disable_connection_drain_on_failover",
				Description: "This can be set to true only if the protocol is TCP  The default is false",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("FailoverPolicy.DisableConnectionDrainOnFailover"),
			},
			{
				Name:        "failover_policy_drop_traffic_if_unhealthy",
				Description: "Applicable only to Failover for Internal TCP/UDP Load Balancing and Network Load Balancing, If set to true, connections to the load balancer are dropped when all primary and all backup backend VMs are unhealthyIf set to false, connections are distributed among all primary VMs when all primary and all backup backend VMs are unhealthy The default is false",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("FailoverPolicy.DropTrafficIfUnhealthy"),
			},
			{
				Name:        "failover_policy_failover_ratio",
				Description: "Applicable only to Failover for Internal TCP/UDP Load Balancing and Network Load Balancing The value of the field must be in the range [0, 1] If the value is 0, the load balancer performs a failover when the number of healthy primary VMs equals zero For all other values, the load balancer performs a failover when the total number of healthy primary VMs is less than this ratio",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("FailoverPolicy.FailoverRatio"),
			},
			{
				Name:        "fingerprint",
				Description: "Fingerprint of this resource A hash of the contents stored in this object This field is used in optimistic locking This field will be ignored when inserting a BackendService An up-to-date fingerprint must be provided in order to update the BackendService, otherwise the request will fail with error 412 conditionNotMet  To see the latest fingerprint, make a get() request to retrieve a BackendService",
				Type:        schema.TypeString,
			},
			{
				Name:        "health_checks",
				Description: "The list of URLs to the healthChecks, httpHealthChecks (legacy), or httpsHealthChecks (legacy) resource for health checking this backend service Not all backend services support legacy health checks See  Load balancer guide Currently, at most one health check can be specified for each backend service Backend services with instance group or zonal NEG backends must have a health check Backend services with internet or serverless NEG backends must not have a health check",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "iap_enabled",
				Description: "Whether the serving infrastructure will authenticate and authorize all incoming requests If true, the oauth2ClientId and oauth2ClientSecret fields must be non-empty",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Iap.Enabled"),
			},
			{
				Name:        "iap_oauth2_client_id",
				Description: "OAuth2 client ID to use for the authentication flow",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Iap.Oauth2ClientId"),
			},
			{
				Name:        "iap_oauth2_client_secret",
				Description: "OAuth2 client secret to use for the authentication flow For security reasons, this value cannot be retrieved via the API Instead, the SHA-256 hash of the value is returned in the oauth2ClientSecretSha256 field",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Iap.Oauth2ClientSecret"),
			},
			{
				Name:        "iap_oauth2_client_secret_sha256",
				Description: "SHA256 hash value for the field oauth2_client_secret above",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Iap.Oauth2ClientSecretSha256"),
			},
			{
				Name:        "id",
				Description: "The unique identifier for the resource This identifier is defined by the server",
				Type:        schema.TypeString,
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "kind",
				Description: "Type of resource Always compute#backendService for backend services",
				Type:        schema.TypeString,
			},
			{
				Name:        "load_balancing_scheme",
				Description: "Specifies the load balancer type Choose EXTERNAL for external HTTP(S), SSL Proxy, TCP Proxy and Network Load Balancing Choose  INTERNAL for Internal TCP/UDP Load Balancing Choose  INTERNAL_MANAGED for Internal HTTP(S) Load Balancing INTERNAL_SELF_MANAGED for Traffic Director A backend service created for one type of load balancer cannot be used with another For more information, refer to Choosing a load balancer",
				Type:        schema.TypeString,
			},
			{
				Name:        "locality_lb_policy",
				Description: "The load balancing algorithm used within the scope of the locality The possible values are: - ROUND_ROBIN: This is a simple policy in which each healthy backend is selected in round robin order This is the default - LEAST_REQUEST: An O(1) algorithm which selects two random healthy hosts and picks the host which has fewer active requests - RING_HASH: The ring/modulo hash load balancer implements consistent hashing to backends The algorithm has the property that the addition/removal of a host from a set of N hosts only affects 1/N of the requests - RANDOM: The load balancer selects a random healthy host - ORIGINAL_DESTINATION: Backend host is selected based on the client connection metadata, ie, connections are opened to the same address as the destination address of the incoming connection before the connection was redirected to the load balancer - MAGLEV: used as a drop in replacement for the ring hash load balancer Maglev is not as stable as ring hash but has faster table lookup build times and host selection times For more information about Maglev, see https://aigoogle/research/pubs/pub44824  This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED  - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED  If sessionAffinity is not NONE, and this field is not set to MAGLEV or RING_HASH, session affinity settings will not take effect  Only the default ROUND_ROBIN policy is supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true",
				Type:        schema.TypeString,
			},
			{
				Name:        "log_config_enable",
				Description: "This field denotes whether to enable logging for the load balancer traffic served by this backend service",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("LogConfig.Enable"),
			},
			{
				Name:        "log_config_sample_rate",
				Description: "This field can only be specified if logging is enabled for this backend service The value of the field must be in [0, 1] This configures the sampling rate of requests to the load balancer where 10 means all logged requests are reported and 00 means no logged requests are reported The default value is 10",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("LogConfig.SampleRate"),
			},
			{
				Name:        "max_stream_duration_nanos",
				Description: "Span of time that's a fraction of a second at nanosecond resolution Durations less than one second are represented with a 0 `seconds` field and a positive `nanos` field Must be from 0 to 999,999,999 inclusive",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("MaxStreamDuration.Nanos"),
			},
			{
				Name:        "max_stream_duration_seconds",
				Description: "Span of time at a resolution of a second Must be from 0 to 315,576,000,000 inclusive Note: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("MaxStreamDuration.Seconds"),
			},
			{
				Name:        "name",
				Description: "Name of the resource Provided by the client when the resource is created The name must be 1-63 characters long, and comply with RFC1035 Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash",
				Type:        schema.TypeString,
			},
			{
				Name:        "network",
				Description: "The URL of the network to which this backend service belongs This field can only be specified when the load balancing scheme is set to INTERNAL",
				Type:        schema.TypeString,
			},
			{
				Name:        "outlier_detection_base_ejection_time_nanos",
				Description: "Span of time that's a fraction of a second at nanosecond resolution Durations less than one second are represented with a 0 `seconds` field and a positive `nanos` field Must be from 0 to 999,999,999 inclusive",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("OutlierDetection.BaseEjectionTime.Nanos"),
			},
			{
				Name:        "outlier_detection_base_ejection_time_seconds",
				Description: "Span of time at a resolution of a second Must be from 0 to 315,576,000,000 inclusive Note: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("OutlierDetection.BaseEjectionTime.Seconds"),
			},
			{
				Name:        "outlier_detection_consecutive_errors",
				Description: "Number of errors before a host is ejected from the connection pool When the backend host is accessed over HTTP, a 5xx return code qualifies as an error Defaults to 5",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("OutlierDetection.ConsecutiveErrors"),
			},
			{
				Name:        "outlier_detection_consecutive_gateway_failure",
				Description: "The number of consecutive gateway failures (502, 503, 504 status or connection errors that are mapped to one of those status codes) before a consecutive gateway failure ejection occurs Defaults to 3",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("OutlierDetection.ConsecutiveGatewayFailure"),
			},
			{
				Name:        "outlier_detection_enforcing_consecutive_errors",
				Description: "The percentage chance that a host will be actually ejected when an outlier status is detected through consecutive 5xx This setting can be used to disable ejection or to ramp it up slowly Defaults to 0",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("OutlierDetection.EnforcingConsecutiveErrors"),
			},
			{
				Name:        "outlier_detection_enforcing_consecutive_gateway_failure",
				Description: "The percentage chance that a host will be actually ejected when an outlier status is detected through consecutive gateway failures This setting can be used to disable ejection or to ramp it up slowly Defaults to 100",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("OutlierDetection.EnforcingConsecutiveGatewayFailure"),
			},
			{
				Name:        "outlier_detection_enforcing_success_rate",
				Description: "The percentage chance that a host will be actually ejected when an outlier status is detected through success rate statistics This setting can be used to disable ejection or to ramp it up slowly Defaults to 100",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("OutlierDetection.EnforcingSuccessRate"),
			},
			{
				Name:        "outlier_detection_interval_nanos",
				Description: "Span of time that's a fraction of a second at nanosecond resolution Durations less than one second are represented with a 0 `seconds` field and a positive `nanos` field Must be from 0 to 999,999,999 inclusive",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("OutlierDetection.Interval.Nanos"),
			},
			{
				Name:        "outlier_detection_interval_seconds",
				Description: "Span of time at a resolution of a second Must be from 0 to 315,576,000,000 inclusive Note: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("OutlierDetection.Interval.Seconds"),
			},
			{
				Name:        "outlier_detection_max_ejection_percent",
				Description: "Maximum percentage of hosts in the load balancing pool for the backend service that can be ejected Defaults to 50%",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("OutlierDetection.MaxEjectionPercent"),
			},
			{
				Name:        "outlier_detection_success_rate_minimum_hosts",
				Description: "The number of hosts in a cluster that must have enough request volume to detect success rate outliers If the number of hosts is less than this setting, outlier detection via success rate statistics is not performed for any host in the cluster Defaults to 5",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("OutlierDetection.SuccessRateMinimumHosts"),
			},
			{
				Name:        "outlier_detection_success_rate_request_volume",
				Description: "The minimum number of total requests that must be collected in one interval (as defined by the interval duration above) to include this host in success rate based outlier detection If the volume is lower than this setting, outlier detection via success rate statistics is not performed for that host Defaults to 100",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("OutlierDetection.SuccessRateRequestVolume"),
			},
			{
				Name:        "outlier_detection_success_rate_stdev_factor",
				Description: "This factor is used to determine the ejection threshold for success rate outlier ejection The ejection threshold is the difference between the mean success rate, and the product of this factor and the standard deviation of the mean success rate: mean - (stdev * success_rate_stdev_factor) This factor is divided by a thousand to get a double That is, if the desired factor is 19, the runtime value should be 1900 Defaults to 1900",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("OutlierDetection.SuccessRateStdevFactor"),
			},
			{
				Name:        "port",
				Description: "Deprecated in favor of portName The TCP port to connect on the backend The default value is 80  Backend services for Internal TCP/UDP Load Balancing and Network Load Balancing require you omit port",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "port_name",
				Description: "A named port on a backend instance group representing the port for communication to the backend VMs in that group Required when the loadBalancingScheme is EXTERNAL (except Network Load Balancing), INTERNAL_MANAGED, or  INTERNAL_SELF_MANAGED and the backends are instance groups The named port must be defined on each backend instance group This parameter has no meaning if the backends are NEGs  Backend services for Internal TCP/UDP Load Balancing and Network Load Balancing require you omit port_name",
				Type:        schema.TypeString,
			},
			{
				Name:        "protocol",
				Description: "The protocol this BackendService uses to communicate with backends  Possible values are HTTP, HTTPS, HTTP2, TCP, SSL, UDP or GRPC depending on the chosen load balancer or Traffic Director configuration Refer to the documentation for the load balancer or for Traffic Director for more information  Must be set to GRPC when the backend service is referenced by a URL map that is bound to target gRPC proxy",
				Type:        schema.TypeString,
			},
			{
				Name:        "region",
				Description: "URL of the region where the regional backend service resides This field is not applicable to global backend services You must specify this field as part of the HTTP request URL It is not settable as a field in the request body",
				Type:        schema.TypeString,
			},
			{
				Name:        "security_policy",
				Description: "The resource URL for the security policy associated with this backend service",
				Type:        schema.TypeString,
			},
			{
				Name:        "security_settings_client_tls_policy",
				Description: "A URL referring to a networksecurityClientTlsPolicy resource that describes how clients should authenticate with this service's backends clientTlsPolicy only applies to a global BackendService with the loadBalancingScheme set to INTERNAL_SELF_MANAGED If left blank, communications are not encrypted Note: This field currently has no impact",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SecuritySettings.ClientTlsPolicy"),
			},
			{
				Name:          "security_settings_subject_alt_names",
				Description:   "A list of Subject Alternative Names (SANs) that the client verifies during a mutual TLS handshake with an server/endpoint for this BackendService When the server presents its X509 certificate to the client, the client inspects the certificate's subjectAltName field If the field contains one of the specified values, the communication continues Otherwise, it fails This additional check enables the client to verify that the server is authorized to run the requested service Note that the contents of the server certificate's subjectAltName field are configured by the Public Key Infrastructure which provisions server identities Only applies to a global BackendService with loadBalancingScheme set to INTERNAL_SELF_MANAGED Only applies when BackendService has an attached clientTlsPolicy with clientCertificate (mTLS mode) Note: This field currently has no impact",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("SecuritySettings.SubjectAltNames"),
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "session_affinity",
				Description: "Type of session affinity to use The default is NONE  When the loadBalancingScheme is EXTERNAL: * For Network Load Balancing, the possible values are NONE, CLIENT_IP, CLIENT_IP_PROTO, or  CLIENT_IP_PORT_PROTO * For all other load balancers that use loadBalancingScheme=EXTERNAL, the possible values are NONE, CLIENT_IP, or GENERATED_COOKIE * You can use GENERATED_COOKIE if the protocol is HTTP, HTTP2, or HTTPS  When the loadBalancingScheme is INTERNAL, possible values are NONE, CLIENT_IP, CLIENT_IP_PROTO, or CLIENT_IP_PORT_PROTO  When the loadBalancingScheme is INTERNAL_SELF_MANAGED, or INTERNAL_MANAGED, possible values are NONE, CLIENT_IP, GENERATED_COOKIE, HEADER_FIELD, or HTTP_COOKIE  Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true",
				Type:        schema.TypeString,
			},
			{
				Name:        "timeout_sec",
				Description: "The backend service timeout has a different meaning depending on the type of load balancer For more information see, Backend service settings The default is 30 seconds The full range of timeout values allowed is 1 - 2,147,483,647 seconds",
				Type:        schema.TypeBigInt,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_compute_backend_service_backends",
				Description: "Message containing information of one individual backend",
				Resolver:    fetchComputeBackendServiceBackends,
				Columns: []schema.Column{
					{
						Name:        "backend_service_cq_id",
						Description: "Unique ID of gcp_compute_backend_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "backend_service_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "balancing_mode",
						Description: "Specifies how to determine whether the backend of a load balancer can handle additional traffic or is fully loaded For usage guidelines, see  Connection balancing mode",
						Type:        schema.TypeString,
					},
					{
						Name:        "capacity_scaler",
						Description: "A multiplier applied to the backend's target capacity of its balancing mode The default value is 1, which means the group serves up to 100% of its configured capacity (depending on balancingMode) A setting of 0 means the group is completely drained, offering 0% of its available capacity The valid ranges are 00 and [01,10] You cannot configure a setting larger than 0 and smaller than 01 You cannot configure a setting of 0 when there is only one backend attached to the backend service",
						Type:        schema.TypeFloat,
					},
					{
						Name:        "description",
						Description: "An optional description of this resource Provide this property when you create the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "failover",
						Description: "This field designates whether this is a failover backend More than one failover backend can be configured for a given BackendService",
						Type:        schema.TypeBool,
					},
					{
						Name:        "group",
						Description: "The fully-qualified URL of an instance group or network endpoint group (NEG) resource The type of backend that a backend service supports depends on",
						Type:        schema.TypeString,
					},
					{
						Name:        "max_connections",
						Description: "Defines a target maximum number of simultaneous connections For usage guidelines, see Connection balancing mode and Utilization balancing mode Not available if the backend's balancingMode is RATE",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "max_connections_per_endpoint",
						Description: "Defines a target maximum number of simultaneous connections For usage guidelines, see Connection balancing mode and Utilization balancing mode  Not available if the backend's balancingMode is RATE Not supported by:  - Internal TCP/UDP Load Balancing - Network Load Balancing",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "max_connections_per_instance",
						Description: "Defines a target maximum number of simultaneous connections For usage guidelines, see Connection balancing mode and Utilization balancing mode  Not available if the backend's balancingMode is RATE Not supported by:  - Internal TCP/UDP Load Balancing - Network Load Balancing",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "max_rate",
						Description: "Defines a maximum number of HTTP requests per second (RPS) For usage guidelines, see Rate balancing mode and Utilization balancing mode  Not available if the backend's balancingMode is CONNECTION",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "max_rate_per_endpoint",
						Description: "Defines a maximum target for requests per second (RPS) For usage guidelines, see Rate balancing mode and Utilization balancing mode  Not available if the backend's balancingMode is CONNECTION",
						Type:        schema.TypeFloat,
					},
					{
						Name:        "max_rate_per_instance",
						Description: "Defines a maximum target for requests per second (RPS) For usage guidelines, see Rate balancing mode and Utilization balancing mode  Not available if the backend's balancingMode is CONNECTION",
						Type:        schema.TypeFloat,
					},
					{
						Name: "max_utilization",
						Type: schema.TypeFloat,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeBackendServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	nextPageToken := ""
	c := meta.(*client.Client)
	for {
		output, err := c.Services.Compute.BackendServices.AggregatedList(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
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
func resolveComputeBackendServiceCdnPolicyBypassCacheOnRequestHeaders(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.BackendService)
	if r.CdnPolicy == nil {
		return nil
	}
	headers := make([]string, len(r.CdnPolicy.BypassCacheOnRequestHeaders))
	for i, v := range r.CdnPolicy.BypassCacheOnRequestHeaders {
		headers[i] = v.HeaderName
	}
	return errors.WithStack(resource.Set("cdn_policy_bypass_cache_on_request_headers", headers))
}
func resolveComputeBackendServiceCdnPolicyNegativeCachingPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.BackendService)
	if r.CdnPolicy == nil {
		return nil
	}

	data, err := json.Marshal(r.CdnPolicy.NegativeCachingPolicy)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(resource.Set(c.Name, data))
}

func fetchComputeBackendServiceBackends(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*compute.BackendService)
	res <- r.Backends
	return nil
}
