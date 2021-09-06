
# Table: gcp_compute_url_maps
Represents a URL Map resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text||
|creation_timestamp|text|Creation timestamp in RFC3339 text format|
|cors_policy_allow_credentials|boolean|In response to a preflight request, setting this to true indicates that the actual request can include user credentials This translates to the Access-Control-Allow-Credentials header Default is false|
|cors_policy_allow_headers|text[]|Specifies the content for the Access-Control-Allow-Headers header|
|cors_policy_allow_methods|text[]|Specifies the content for the Access-Control-Allow-Methods header|
|cors_policy_allow_origin_regexes|text[]|Specifies the regualar expression patterns that match allowed origins For regular expression grammar please see githubcom/google/re2/wiki/Syntax An origin is allowed if it matches either an item in allowOrigins or an item in allowOriginRegexes|
|cors_policy_allow_origins|text[]|Specifies the list of origins that will be allowed to do CORS requests An origin is allowed if it matches either an item in allowOrigins or an item in allowOriginRegexes|
|cors_policy_disabled|boolean|If true, specifies the CORS policy is disabled The default value of false, which indicates that the CORS policy is in effect|
|cors_policy_expose_headers|text[]|Specifies the content for the Access-Control-Expose-Headers header|
|cors_policy_max_age|bigint|Specifies how long results of a preflight request can be cached in seconds This translates to the Access-Control-Max-Age header|
|fault_injection_policy_abort_http_status|bigint|The HTTP status code used to abort the request The value must be between 200 and 599 inclusive|
|fault_injection_policy_abort_percentage|float|The percentage of traffic (connections/operations/requests) which will be aborted as part of fault injection The value must be between 00 and 1000 inclusive|
|fault_injection_policy_delay_fixed_delay_nanos|bigint|Span of time that's a fraction of a second at nanosecond resolution Durations less than one second are represented with a 0 `seconds` field and a positive `nanos` field Must be from 0 to 999,999,999 inclusive|
|fault_injection_policy_delay_fixed_delay_seconds|bigint|Span of time at a resolution of a second Must be from 0 to 315,576,000,000 inclusive Note: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365|
|fault_injection_policy_delay_percentage|float|The percentage of traffic (connections/operations/requests) on which delay will be introduced as part of fault injection The value must be between 00 and 1000 inclusive|
|max_stream_duration_nanos|bigint|Span of time that's a fraction of a second at nanosecond resolution Durations less than one second are represented with a 0 `seconds` field and a positive `nanos` field Must be from 0 to 999,999,999 inclusive|
|max_stream_duration_seconds|bigint|Span of time at a resolution of a second Must be from 0 to 315,576,000,000 inclusive Note: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365|
|request_mirror_policy_backend_service|text|The full or partial URL to the BackendService resource being mirrored to|
|retry_policy_num_retries|bigint|Specifies the allowed number retries This number must be > 0 If not specified, defaults to 1|
|retry_policy_per_try_timeout_nanos|bigint|Span of time that's a fraction of a second at nanosecond resolution Durations less than one second are represented with a 0 `seconds` field and a positive `nanos` field Must be from 0 to 999,999,999 inclusive|
|retry_policy_per_try_timeout_seconds|bigint|Span of time at a resolution of a second Must be from 0 to 315,576,000,000 inclusive Note: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365|
|retry_policy_retry_conditions|text[]|Specfies one or more conditions when this retry rule applies|
|timeout_nanos|bigint|Span of time that's a fraction of a second at nanosecond resolution Durations less than one second are represented with a 0 `seconds` field and a positive `nanos` field Must be from 0 to 999,999,999 inclusive|
|timeout_seconds|bigint|Span of time at a resolution of a second Must be from 0 to 315,576,000,000 inclusive Note: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365|
|url_rewrite_host_rewrite|text|Prior to forwarding the request to the selected service, the request's host header is replaced with contents of hostRewrite The value must be between 1 and 255 characters|
|url_rewrite_path_prefix_rewrite|text|Prior to forwarding the request to the selected backend service, the matching portion of the request's path is replaced by pathPrefixRewrite The value must be between 1 and 1024 characters|
|default_service|text|The full or partial URL of the defaultService resource to which traffic is directed|
|default_url_redirect_host_redirect|text|The host that will be used in the redirect response instead of the one that was supplied in the request The value must be between 1 and 255 characters|
|default_url_redirect_https_redirect|boolean|If set to true, the URL scheme in the redirected request is set to https If set to false, the URL scheme of the redirected request will remain the same as that of the request|
|default_url_redirect_path_redirect|text|The path that will be used in the redirect response instead of the one that was supplied in the request|
|default_url_redirect_prefix_redirect|text|The prefix that replaces the prefixMatch specified in the HttpRouteRuleMatch, retaining the remaining portion of the URL before redirecting the request prefixRedirect cannot be supplied together with pathRedirect Supply one alone or neither If neither is supplied, the path of the original request will be used for the redirect The value must be between 1 and 1024 characters|
|default_url_redirect_redirect_response_code|text|The HTTP Status code to use for this RedirectAction Supported values are: - MOVED_PERMANENTLY_DEFAULT, which is the default value and corresponds to 301 - FOUND, which corresponds to 302 - SEE_OTHER which corresponds to 303 - TEMPORARY_REDIRECT, which corresponds to 307 In this case, the request method will be retained - PERMANENT_REDIRECT, which corresponds to 308 In this case, the request method will be retained|
|default_url_redirect_strip_query|boolean|If set to true, any accompanying query portion of the original URL is removed prior to redirecting the request If set to false, the query portion of the original URL is retained The default is set to false|
|description|text|An optional description of this resource Provide this property when you create the resource|
|fingerprint|text|Fingerprint of this resource A hash of the contents stored in this object This field is used in optimistic locking This field will be ignored when inserting a UrlMap An up-to-date fingerprint must be provided in order to update the UrlMap, otherwise the request will fail with error 412 conditionNotMet  To see the latest fingerprint, make a get() request to retrieve a UrlMap|
|header_action_request_headers_to_add|jsonb|Headers to add to a matching request prior to forwarding the request to the backendService|
|header_action_request_headers_to_remove|text[]|A list of header names for headers that need to be removed from the request prior to forwarding the request to the backendService|
|header_action_response_headers_to_add|jsonb|Headers to add the response prior to sending the response back to the client|
|header_action_response_headers_to_remove|text[]|A list of header names for headers that need to be removed from the response prior to sending the response back to the client|
|id|bigint|The unique identifier for the resource This identifier is defined by the server|
|kind|text|Type of the resource Always compute#urlMaps for url maps|
|name|text|Name of the resource Provided by the client when the resource is created The name must be 1-63 characters long, and comply with RFC1035 Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash|
|region|text|URL of the region where the regional URL map resides This field is not applicable to global URL maps You must specify this field as part of the HTTP request URL It is not settable as a field in the request body|
|self_link|text|Server-defined URL for the resource|
