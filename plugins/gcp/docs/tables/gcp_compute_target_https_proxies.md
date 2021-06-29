
# Table: gcp_compute_target_https_proxies
Represents a Target HTTPS Proxy resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|authorization_policy|text|A URL referring to a networksecurityAuthorizationPolicy resource that describes how the proxy should authorize inbound traffic If left blank, access will not be restricted by an authorization policy Refer to the AuthorizationPolicy resource for additional details authorizationPolicy only applies to a global TargetHttpsProxy attached to globalForwardingRules with the loadBalancingScheme set to INTERNAL_SELF_MANAGED Note: This field currently has no impact|
|creation_timestamp|timestamp without time zone|Creation timestamp in RFC3339 text format|
|description|text|An optional description of this resource|
|fingerprint|text|Fingerprint of this resource|
|https_proxy_id|text|Unique Id of the ssl proxy|
|kind|text|Type of resource Always compute#targetHttpsProxy for target HTTPS proxies|
|name|text|Name of the resource|
|proxy_bind|boolean|This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED  When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy) The Envoy proxy listens for inbound requests and handles requests when it receives them  The default is false|
|quic_override|text|Specifies the QUIC override policy for this TargetHttpsProxy resource This setting determines whether the load balancer attempts to negotiate QUIC with clients You can specify NONE, ENABLE, or DISABLE - When quic-override is set to NONE, Google manages whether QUIC is used - When quic-override is set to ENABLE, the load balancer uses QUIC when possible - When quic-override is set to DISABLE, the load balancer doesn't use QUIC - If the quic-override flag is not specified, NONE is implied|
|region|text|URL of the region where the regional TargetHttpsProxy resides This field is not applicable to global TargetHttpsProxies|
|self_link|text|Server-defined URL for the resource|
|server_tls_policy|text|A URL referring to a networksecurityServerTlsPolicy resource that describes how the proxy should authenticate inbound traffic serverTlsPolicy only applies to a global TargetHttpsProxy attached to globalForwardingRules with the loadBalancingScheme set to INTERNAL_SELF_MANAGED If left blank, communications are not encrypted Note: This field currently has no impact|
|ssl_certificates|text[]|URLs to SslCertificate resources that are used to authenticate connections between users and the load balancer At least one SSL certificate must be specified Currently, you may specify up to 15 SSL certificates|
|ssl_policy|text|URL of SslPolicy resource that will be associated with the TargetHttpsProxy resource If not set, the TargetHttpsProxy resource has no SSL policy configured|
|url_map|text|A fully-qualified or valid partial URL to the UrlMap resource that defines the mapping from URL to the BackendService For example, the following are all valid URLs for specifying a URL map: - https://wwwgoogleapis|
