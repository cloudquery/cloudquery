package compute

import (
	"context"
	"encoding/json"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/compute/v1"
)

func ComputeURLMaps() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_url_maps",
		Description: "Represents a URL Map resource",
		Resolver:    fetchComputeUrlMaps,
		Multiplex:   client.ProjectMultiplex,

		Options: schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "id"}},
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:        "creation_timestamp",
				Description: "Creation timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "cors_policy_allow_credentials",
				Description: "In response to a preflight request, setting this to true indicates that the actual request can include user credentials This translates to the Access-Control-Allow-Credentials header Default is false",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DefaultRouteAction.CorsPolicy.AllowCredentials"),
			},
			{
				Name:          "cors_policy_allow_headers",
				Description:   "Specifies the content for the Access-Control-Allow-Headers header",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("DefaultRouteAction.CorsPolicy.AllowHeaders"),
			},
			{
				Name:          "cors_policy_allow_methods",
				Description:   "Specifies the content for the Access-Control-Allow-Methods header",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("DefaultRouteAction.CorsPolicy.AllowMethods"),
			},
			{
				Name:          "cors_policy_allow_origin_regexes",
				Description:   "Specifies the regualar expression patterns that match allowed origins For regular expression grammar please see githubcom/google/re2/wiki/Syntax An origin is allowed if it matches either an item in allowOrigins or an item in allowOriginRegexes",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("DefaultRouteAction.CorsPolicy.AllowOriginRegexes"),
			},
			{
				Name:          "cors_policy_allow_origins",
				Description:   "Specifies the list of origins that will be allowed to do CORS requests An origin is allowed if it matches either an item in allowOrigins or an item in allowOriginRegexes",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("DefaultRouteAction.CorsPolicy.AllowOrigins"),
			},
			{
				Name:        "cors_policy_disabled",
				Description: "If true, specifies the CORS policy is disabled The default value of false, which indicates that the CORS policy is in effect",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DefaultRouteAction.CorsPolicy.Disabled"),
			},
			{
				Name:          "cors_policy_expose_headers",
				Description:   "Specifies the content for the Access-Control-Expose-Headers header",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("DefaultRouteAction.CorsPolicy.ExposeHeaders"),
			},
			{
				Name:        "cors_policy_max_age",
				Description: "Specifies how long results of a preflight request can be cached in seconds This translates to the Access-Control-Max-Age header",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DefaultRouteAction.CorsPolicy.MaxAge"),
			},
			{
				Name:        "fault_injection_policy_abort_http_status",
				Description: "The HTTP status code used to abort the request The value must be between 200 and 599 inclusive",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DefaultRouteAction.FaultInjectionPolicy.Abort.HttpStatus"),
			},
			{
				Name:        "fault_injection_policy_abort_percentage",
				Description: "The percentage of traffic (connections/operations/requests) which will be aborted as part of fault injection The value must be between 00 and 1000 inclusive",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("DefaultRouteAction.FaultInjectionPolicy.Abort.Percentage"),
			},
			{
				Name:        "fault_injection_policy_delay_fixed_delay_nanos",
				Description: "Span of time that's a fraction of a second at nanosecond resolution Durations less than one second are represented with a 0 `seconds` field and a positive `nanos` field Must be from 0 to 999,999,999 inclusive",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DefaultRouteAction.FaultInjectionPolicy.Delay.FixedDelay.Nanos"),
			},
			{
				Name:        "fault_injection_policy_delay_fixed_delay_seconds",
				Description: "Span of time at a resolution of a second Must be from 0 to 315,576,000,000 inclusive Note: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DefaultRouteAction.FaultInjectionPolicy.Delay.FixedDelay.Seconds"),
			},
			{
				Name:        "fault_injection_policy_delay_percentage",
				Description: "The percentage of traffic (connections/operations/requests) on which delay will be introduced as part of fault injection The value must be between 00 and 1000 inclusive",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("DefaultRouteAction.FaultInjectionPolicy.Delay.Percentage"),
			},
			{
				Name:        "max_stream_duration_nanos",
				Description: "Span of time that's a fraction of a second at nanosecond resolution Durations less than one second are represented with a 0 `seconds` field and a positive `nanos` field Must be from 0 to 999,999,999 inclusive",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DefaultRouteAction.MaxStreamDuration.Nanos"),
			},
			{
				Name:        "max_stream_duration_seconds",
				Description: "Span of time at a resolution of a second Must be from 0 to 315,576,000,000 inclusive Note: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DefaultRouteAction.MaxStreamDuration.Seconds"),
			},
			{
				Name:        "request_mirror_policy_backend_service",
				Description: "The full or partial URL to the BackendService resource being mirrored to",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DefaultRouteAction.RequestMirrorPolicy.BackendService"),
			},
			{
				Name:        "retry_policy_num_retries",
				Description: "Specifies the allowed number retries This number must be > 0 If not specified, defaults to 1",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DefaultRouteAction.RetryPolicy.NumRetries"),
			},
			{
				Name:        "retry_policy_per_try_timeout_nanos",
				Description: "Span of time that's a fraction of a second at nanosecond resolution Durations less than one second are represented with a 0 `seconds` field and a positive `nanos` field Must be from 0 to 999,999,999 inclusive",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DefaultRouteAction.RetryPolicy.PerTryTimeout.Nanos"),
			},
			{
				Name:        "retry_policy_per_try_timeout_seconds",
				Description: "Span of time at a resolution of a second Must be from 0 to 315,576,000,000 inclusive Note: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DefaultRouteAction.RetryPolicy.PerTryTimeout.Seconds"),
			},
			{
				Name:          "retry_policy_retry_conditions",
				Description:   "Specfies one or more conditions when this retry rule applies",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("DefaultRouteAction.RetryPolicy.RetryConditions"),
			},
			{
				Name:        "timeout_nanos",
				Description: "Span of time that's a fraction of a second at nanosecond resolution Durations less than one second are represented with a 0 `seconds` field and a positive `nanos` field Must be from 0 to 999,999,999 inclusive",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DefaultRouteAction.Timeout.Nanos"),
			},
			{
				Name:        "timeout_seconds",
				Description: "Span of time at a resolution of a second Must be from 0 to 315,576,000,000 inclusive Note: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DefaultRouteAction.Timeout.Seconds"),
			},
			{
				Name:        "url_rewrite_host_rewrite",
				Description: "Prior to forwarding the request to the selected service, the request's host header is replaced with contents of hostRewrite The value must be between 1 and 255 characters",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DefaultRouteAction.UrlRewrite.HostRewrite"),
			},
			{
				Name:        "url_rewrite_path_prefix_rewrite",
				Description: "Prior to forwarding the request to the selected backend service, the matching portion of the request's path is replaced by pathPrefixRewrite The value must be between 1 and 1024 characters",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DefaultRouteAction.UrlRewrite.PathPrefixRewrite"),
			},
			{
				Name:        "default_service",
				Description: "The full or partial URL of the defaultService resource to which traffic is directed",
				Type:        schema.TypeString,
			},
			{
				Name:        "default_url_redirect_host_redirect",
				Description: "The host that will be used in the redirect response instead of the one that was supplied in the request The value must be between 1 and 255 characters",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DefaultUrlRedirect.HostRedirect"),
			},
			{
				Name:        "default_url_redirect_https_redirect",
				Description: "If set to true, the URL scheme in the redirected request is set to https If set to false, the URL scheme of the redirected request will remain the same as that of the request",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DefaultUrlRedirect.HttpsRedirect"),
			},
			{
				Name:        "default_url_redirect_path_redirect",
				Description: "The path that will be used in the redirect response instead of the one that was supplied in the request",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DefaultUrlRedirect.PathRedirect"),
			},
			{
				Name:        "default_url_redirect_prefix_redirect",
				Description: "The prefix that replaces the prefixMatch specified in the HttpRouteRuleMatch, retaining the remaining portion of the URL before redirecting the request prefixRedirect cannot be supplied together with pathRedirect Supply one alone or neither If neither is supplied, the path of the original request will be used for the redirect The value must be between 1 and 1024 characters",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DefaultUrlRedirect.PrefixRedirect"),
			},
			{
				Name:        "default_url_redirect_redirect_response_code",
				Description: "The HTTP Status code to use for this RedirectAction Supported values are: - MOVED_PERMANENTLY_DEFAULT, which is the default value and corresponds to 301 - FOUND, which corresponds to 302 - SEE_OTHER which corresponds to 303 - TEMPORARY_REDIRECT, which corresponds to 307 In this case, the request method will be retained - PERMANENT_REDIRECT, which corresponds to 308 In this case, the request method will be retained",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DefaultUrlRedirect.RedirectResponseCode"),
			},
			{
				Name:        "default_url_redirect_strip_query",
				Description: "If set to true, any accompanying query portion of the original URL is removed prior to redirecting the request If set to false, the query portion of the original URL is retained The default is set to false",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DefaultUrlRedirect.StripQuery"),
			},
			{
				Name:        "description",
				Description: "An optional description of this resource Provide this property when you create the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "fingerprint",
				Description: "Fingerprint of this resource A hash of the contents stored in this object This field is used in optimistic locking This field will be ignored when inserting a UrlMap An up-to-date fingerprint must be provided in order to update the UrlMap, otherwise the request will fail with error 412 conditionNotMet  To see the latest fingerprint, make a get() request to retrieve a UrlMap",
				Type:        schema.TypeString,
			},
			{
				Name:          "header_action_request_headers_to_add",
				Description:   "Headers to add to a matching request prior to forwarding the request to the backendService",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
				Resolver:      resolveComputeURLMapHeaderActionRequestHeadersToAdd,
			},
			{
				Name:          "header_action_request_headers_to_remove",
				Description:   "A list of header names for headers that need to be removed from the request prior to forwarding the request to the backendService",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("HeaderAction.RequestHeadersToRemove"),
			},
			{
				Name:          "header_action_response_headers_to_add",
				Description:   "Headers to add the response prior to sending the response back to the client",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
				Resolver:      resolveComputeURLMapHeaderActionResponseHeadersToAdd,
			},
			{
				Name:          "header_action_response_headers_to_remove",
				Description:   "A list of header names for headers that need to be removed from the response prior to sending the response back to the client",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("HeaderAction.ResponseHeadersToRemove"),
			},
			{
				Name:        "id",
				Description: "The unique identifier for the resource This identifier is defined by the server",
				Type:        schema.TypeString,
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "kind",
				Description: "Type of the resource Always compute#urlMaps for url maps",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "Name of the resource Provided by the client when the resource is created The name must be 1-63 characters long, and comply with RFC1035 Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash",
				Type:        schema.TypeString,
			},
			{
				Name:        "region",
				Description: "URL of the region where the regional URL map resides This field is not applicable to global URL maps You must specify this field as part of the HTTP request URL It is not settable as a field in the request body",
				Type:        schema.TypeString,
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "gcp_compute_url_map_weighted_backend_services",
				Description:   "In contrast to a single BackendService in HttpRouteAction to which all matching traffic is directed to, WeightedBackendService allows traffic to be split across multiple BackendServices",
				Resolver:      fetchComputeUrlMapWeightedBackendServices,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "url_map_cq_id",
						Description: "Unique CloudQuery ID of gcp_compute_url_maps table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "backend_service",
						Description: "The full or partial URL to the default BackendService resource Before forwarding the request to backendService, the loadbalancer applies any relevant headerActions specified as part of this backendServiceWeight",
						Type:        schema.TypeString,
					},
					{
						Name:        "header_action",
						Description: "Specifies changes to request and response headers that need to take effect for the selected backendService headerAction specified here take effect before headerAction in the enclosing HttpRouteRule, PathMatcher and UrlMap Note that headerAction is not supported for Loadbalancers that have their loadBalancingScheme set to EXTERNAL Not supported when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true",
						Type:        schema.TypeJSON,
						Resolver:    resolveComputeURLMapWeightedBackendServiceHeaderAction,
					},
					{
						Name:        "weight",
						Description: "Specifies the fraction of traffic sent to backendService, computed as weight / (sum of all weightedBackendService weights in routeAction)  The selection of a backend service is determined only for new traffic Once a user's request has been directed to a backendService, subsequent requests will be sent to the same backendService as determined by the BackendService's session affinity policy",
						Type:        schema.TypeBigInt,
					},
				},
			},
			{
				Name:        "gcp_compute_url_map_host_rules",
				Description: "UrlMaps A host-matching rule for a URL If matched, will use the named PathMatcher to select the BackendService",
				Resolver:    fetchComputeUrlMapHostRules,
				Columns: []schema.Column{
					{
						Name:        "url_map_cq_id",
						Description: "Unique CloudQuery ID of gcp_compute_url_maps table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "description",
						Description: "An optional description of this resource Provide this property when you create the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "hosts",
						Description: "The list of host patterns to match They must be valid hostnames with optional port numbers in the format host:port * matches any string of ([a-z0-9-]*) In that case, * must be the first character and must be followed in the pattern by either - or  * based matching is not supported when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "path_matcher",
						Description: "The name of the PathMatcher to use to match the path portion of the URL if the hostRule matches the URL's host portion",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "gcp_compute_url_map_path_matchers",
				Description: "A matcher for the path portion of the URL The BackendService from the longest-matched rule will serve the URL If no rule was matched, the default service will be used",
				Resolver:    fetchComputeUrlMapPathMatchers,
				Columns: []schema.Column{
					{
						Name:        "url_map_cq_id",
						Description: "Unique CloudQuery ID of gcp_compute_url_maps table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:          "default_route_action",
						Description:   "defaultRouteAction takes effect when none of the pathRules or routeRules match The load balancer performs advanced routing actions like URL rewrites, header transformations, etc prior to forwarding the request to the selected backend If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set Conversely if defaultService is set, defaultRouteAction cannot contain any  weightedBackendServices Only one of defaultRouteAction or defaultUrlRedirect must be set UrlMaps for external HTTP(S) load balancers support only the urlRewrite action within a pathMatcher's defaultRouteAction",
						Type:          schema.TypeJSON,
						IgnoreInTests: true,
						Resolver:      resolveComputeURLMapPathMatcherDefaultRouteAction,
					},
					{
						Name:        "default_service",
						Description: "The full or partial URL to the BackendService resource This will be used if none of the pathRules or routeRules defined by this PathMatcher are matched For example, the following are all valid URLs to a BackendService resource: - https://wwwgoogleapiscom/compute/v1/projects/project/global/backendServices/backendService  - compute/v1/projects/project/global/backendServices/backendService  - global/backendServices/backendService  If defaultRouteAction is additionally specified, advanced routing actions like URL Rewrites, etc take effect prior to sending the request to the backend However, if defaultService is specified, defaultRouteAction cannot contain any weightedBackendServices Conversely, if defaultRouteAction specifies any weightedBackendServices, defaultService must not be specified Only one of defaultService, defaultUrlRedirect  or defaultRouteActionweightedBackendService must be set Authorization requires one or more of the following Google IAM permissions on the specified resource default_service: - computebackendBucketsuse - computebackendServices",
						Type:        schema.TypeString,
					},
					{
						Name:        "default_url_redirect_host_redirect",
						Description: "The host that will be used in the redirect response instead of the one that was supplied in the request The value must be between 1 and 255 characters",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DefaultUrlRedirect.HostRedirect"),
					},
					{
						Name:        "default_url_redirect_https_redirect",
						Description: "If set to true, the URL scheme in the redirected request is set to https If set to false, the URL scheme of the redirected request will remain the same as that of the request This must only be set for UrlMaps used in TargetHttpProxys Setting this true for TargetHttpsProxy is not permitted The default is set to false",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DefaultUrlRedirect.HttpsRedirect"),
					},
					{
						Name:        "default_url_redirect_path_redirect",
						Description: "The path that will be used in the redirect response instead of the one that was supplied in the request pathRedirect cannot be supplied together with prefixRedirect Supply one alone or neither If neither is supplied, the path of the original request will be used for the redirect The value must be between 1 and 1024 characters",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DefaultUrlRedirect.PathRedirect"),
					},
					{
						Name:        "default_url_redirect_prefix_redirect",
						Description: "The prefix that replaces the prefixMatch specified in the HttpRouteRuleMatch, retaining the remaining portion of the URL before redirecting the request prefixRedirect cannot be supplied together with pathRedirect Supply one alone or neither If neither is supplied, the path of the original request will be used for the redirect The value must be between 1 and 1024 characters",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DefaultUrlRedirect.PrefixRedirect"),
					},
					{
						Name:        "default_url_redirect_redirect_response_code",
						Description: "The HTTP Status code to use for this RedirectAction Supported values are: - MOVED_PERMANENTLY_DEFAULT, which is the default value and corresponds to 301 - FOUND, which corresponds to 302 - SEE_OTHER which corresponds to 303 - TEMPORARY_REDIRECT, which corresponds to 307 In this case, the request method will be retained - PERMANENT_REDIRECT, which corresponds to 308 In this case, the request method will be retained",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DefaultUrlRedirect.RedirectResponseCode"),
					},
					{
						Name:        "default_url_redirect_strip_query",
						Description: "If set to true, any accompanying query portion of the original URL is removed prior to redirecting the request If set to false, the query portion of the original URL is retained The default is set to false",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DefaultUrlRedirect.StripQuery"),
					},
					{
						Name:        "description",
						Description: "An optional description of this resource Provide this property when you create the resource",
						Type:        schema.TypeString,
					},
					{
						Name:          "header_action",
						Description:   "Specifies changes to request and response headers that need to take effect for the selected backendService HeaderAction specified here are applied after the matching HttpRouteRule HeaderAction and before the HeaderAction in the UrlMap  Note that headerAction is not supported for Loadbalancers that have their loadBalancingScheme set to EXTERNAL Not supported when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true",
						Type:          schema.TypeJSON,
						IgnoreInTests: true,
						Resolver:      resolveComputeURLMapPathMatcherHeaderAction,
					},
					{
						Name:        "name",
						Description: "The name to which this PathMatcher is referred by the HostRule",
						Type:        schema.TypeString,
					},
					{
						Name:        "path_rules",
						Description: "The list of path rules Use this list instead of routeRules when routing based on simple path matching is all that's required The order by which path rules are specified does not matter Matches are always done on the longest-path-first basis For example: a pathRule with a path /a/b/c/* will match before /a/b/* irrespective of the order in which those paths appear in this list Within a given pathMatcher, only one of pathRules or routeRules must be set",
						Type:        schema.TypeJSON,
						Resolver:    resolveComputeURLMapPathMatcherPathRules,
					},
					{
						Name:          "route_rules",
						Description:   "The list of HTTP route rules Use this list instead of pathRules when advanced route matching and routing actions are desired routeRules are evaluated in order of priority, from the lowest to highest number Within a given pathMatcher, you can set only one of pathRules or routeRules",
						Type:          schema.TypeJSON,
						IgnoreInTests: true,
						Resolver:      resolveComputeURLMapPathMatcherRouteRules,
					},
				},
			},
			{
				Name:        "gcp_compute_url_map_tests",
				Description: "Message for the expected URL mappings",
				Resolver:    fetchComputeUrlMapTests,
				Columns: []schema.Column{
					{
						Name:        "url_map_cq_id",
						Description: "Unique CloudQuery ID of gcp_compute_url_maps table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "description",
						Description: "Description of this test case",
						Type:        schema.TypeString,
					},
					{
						Name:        "expected_output_url",
						Description: "The expected output URL evaluated by load balancer containing the scheme, host, path and query parameters For rules that forward requests to backends, the test passes only when expectedOutputUrl matches the request forwarded by load balancer to backends For rules with urlRewrite, the test verifies that the forwarded request matches hostRewrite and pathPrefixRewrite in the urlRewrite action When service is specified, expectedOutputUrl`s scheme is ignored For rules with urlRedirect, the test passes only if expectedOutputUrl matches the URL in the load balancer's redirect response If urlRedirect specifies https_redirect, the test passes only if the scheme in expectedOutputUrl is also set to https If urlRedirect specifies strip_query, the test passes only if expectedOutputUrl does not contain any query parameters expectedOutputUrl is optional when service is specified",
						Type:        schema.TypeString,
					},
					{
						Name:        "expected_redirect_response_code",
						Description: "For rules with urlRedirect, the test passes only if expectedRedirectResponseCode matches the HTTP status code in load balancer's redirect response expectedRedirectResponseCode cannot be set when service is set",
						Type:        schema.TypeBigInt,
					},
					{
						Name:          "headers",
						Description:   "HTTP headers for this request If headers contains a host header, then host must also match the header value",
						Type:          schema.TypeJSON,
						IgnoreInTests: true,
						Resolver:      resolveComputeURLMapTestHeaders,
					},
					{
						Name:        "host",
						Description: "Host portion of the URL If headers contains a host header, then host must also match the header value",
						Type:        schema.TypeString,
					},
					{
						Name:        "path",
						Description: "Path portion of the URL",
						Type:        schema.TypeString,
					},
					{
						Name:        "service",
						Description: "Expected BackendService or BackendBucket resource the given URL should be mapped to service cannot be set if expectedRedirectResponseCode is set",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeUrlMaps(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Compute.UrlMaps.List(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		res <- output.Items

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func resolveComputeURLMapHeaderActionRequestHeadersToAdd(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.UrlMap)
	if r.HeaderAction == nil || r.HeaderAction.RequestHeadersToAdd == nil {
		return nil
	}

	var j []interface{}
	data, err := json.Marshal(r.HeaderAction.RequestHeadersToAdd)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := json.Unmarshal(data, &j); err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(resource.Set(c.Name, j))
}
func resolveComputeURLMapHeaderActionResponseHeadersToAdd(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.UrlMap)
	if r.HeaderAction == nil || r.HeaderAction.ResponseHeadersToAdd == nil {
		return nil
	}

	var j []interface{}
	data, err := json.Marshal(r.HeaderAction.ResponseHeadersToAdd)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := json.Unmarshal(data, &j); err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(resource.Set(c.Name, j))
}
func fetchComputeUrlMapWeightedBackendServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*compute.UrlMap)
	if r.DefaultRouteAction == nil || r.DefaultRouteAction.WeightedBackendServices == nil {
		return nil
	}

	res <- r.DefaultRouteAction.WeightedBackendServices
	return nil
}
func resolveComputeURLMapWeightedBackendServiceHeaderAction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.WeightedBackendService)
	var j map[string]interface{}
	data, err := json.Marshal(r.HeaderAction)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := json.Unmarshal(data, &j); err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(resource.Set(c.Name, j))
}
func fetchComputeUrlMapHostRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*compute.UrlMap)
	if r.HostRules == nil {
		return nil
	}

	res <- r.HostRules
	return nil
}
func fetchComputeUrlMapPathMatchers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*compute.UrlMap)
	if r.PathMatchers == nil {
		return nil
	}

	res <- r.PathMatchers
	return nil
}
func resolveComputeURLMapPathMatcherDefaultRouteAction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.PathMatcher)
	var j map[string]interface{}
	data, err := json.Marshal(r.DefaultRouteAction)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := json.Unmarshal(data, &j); err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(resource.Set(c.Name, j))
}
func resolveComputeURLMapPathMatcherHeaderAction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.PathMatcher)
	var j map[string]interface{}
	data, err := json.Marshal(r.HeaderAction)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := json.Unmarshal(data, &j); err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(resource.Set(c.Name, j))
}
func resolveComputeURLMapPathMatcherPathRules(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.PathMatcher)
	var j []interface{}
	data, err := json.Marshal(r.PathRules)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := json.Unmarshal(data, &j); err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(resource.Set(c.Name, j))
}
func resolveComputeURLMapPathMatcherRouteRules(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.PathMatcher)
	var j []interface{}
	data, err := json.Marshal(r.RouteRules)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := json.Unmarshal(data, &j); err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(resource.Set(c.Name, j))
}
func fetchComputeUrlMapTests(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*compute.UrlMap)
	if r.Tests == nil {
		return nil
	}

	res <- r.Tests
	return nil
}
func resolveComputeURLMapTestHeaders(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.UrlMapTest)
	j := make(map[string]interface{})
	if r.Headers == nil {
		return nil
	}

	for _, h := range r.Headers {
		j[h.Name] = h.Value
	}

	return errors.WithStack(resource.Set(c.Name, j))
}
