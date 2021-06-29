
# Table: gcp_compute_ssl_policies
Represents an SSL Policy resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|creation_timestamp|timestamp without time zone|Creation timestamp in RFC3339 text format|
|custom_features|text[]|A list of features enabled when the selected profile is CUSTOM The - method returns the set of features that can be specified in this list This field must be empty if the profile is not CUSTOM|
|description|text|An optional description of this resource|
|enabled_features|text[]|The list of features enabled in the SSL policy|
|fingerprint|text|Fingerprint of this resource|
|ssl_policy_id|text|The unique identifier for the resource This identifier is defined by the server|
|kind|text|Type of the resource Always compute#sslPolicy for SSL policies|
|min_tls_version|text|The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer This can be one of TLS_1_0, TLS_1_1, TLS_1_2|
|name|text|Name of the resource|
|profile|text|Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field|
|self_link|text|Server-defined URL for the resource|
