
# Table: gcp_compute_target_ssl_proxies
Represents a Target SSL Proxy resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|creation_timestamp|timestamp without time zone|Creation timestamp in RFC3339 text format|
|description|text|An optional description of this resource|
|ssl_proxy_id|text|The unique identifier for the resource This identifier is defined by the server|
|kind|text|Type of the resource Always compute#targetSslProxy for target SSL proxies|
|name|text|Name of the resource|
|proxy_header|text|Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1 The default is NONE|
|self_link|text|Server-defined URL for the resource|
|service|text|URL to the BackendService resource|
|ssl_certificates|text[]|URLs to SslCertificate resources that are used to authenticate connections to Backends At least one SSL certificate must be specified Currently, you may specify up to 15 SSL certificates|
|ssl_policy|text|URL of SslPolicy resource that will be associated with the TargetSslProxy resource If not set, the TargetSslProxy resource will not have any SSL policy configured|
