
# Table: gcp_compute_target_http_proxies
Represents a Target HTTP Proxy resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|creation_timestamp|timestamp without time zone|Creation timestamp in RFC3339 text format|
|description|text|An optional description of this resource|
|fingerprint|text|Fingerprint of this resource|
|id|text|The unique identifier for the resource|
|kind|text|Type of resource Always compute#targetHttpProxy for target HTTP proxies|
|name|text|Name of the resource|
|proxy_bind|boolean|This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED  When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy) The Envoy proxy listens for inbound requests and handles requests when it receives them  The default is false|
|region|text|URL of the region where the regional Target HTTP Proxy resides. This field is not applicable to global Target HTTP Proxies.|
|self_link|text|Server-defined URL for the resource|
|url_map|text|A fully-qualified or valid partial URL to the UrlMap resource that defines the mapping from URL to the BackendService.|
