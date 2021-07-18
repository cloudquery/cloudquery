
# Table: gcp_compute_disk_types
Represents a Disk Type resource.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|creation_timestamp|text|Creation timestamp in RFC3339 text format|
|default_disk_size_gb|bigint|Server-defined default disk size in GB|
|deprecated_deleted|text|An optional RFC3339 timestamp on or after which the state of this resource is intended to change to DELETED This is only informational and the status will not change unless the client explicitly changes it|
|deprecated|text||
|deprecated_obsolete|text|An optional RFC3339 timestamp on or after which the state of this resource is intended to change to OBSOLETE This is only informational and the status will not change unless the client explicitly changes it|
|deprecated_replacement|text|The URL of the suggested replacement for a deprecated resource The suggested replacement resource must be the same kind of resource as the deprecated resource|
|deprecated_state|text|The deprecation state of this resource This can be ACTIVE, DEPRECATED, OBSOLETE, or DELETED Operations which communicate the end of life date for an image, can use ACTIVE Operations which create a new resource using a DEPRECATED resource will return successfully, but with a warning indicating the deprecated resource and recommending its replacement Operations which use OBSOLETE or DELETED resources will be rejected and result in an error|
|description|text|An optional description of this resource|
|id|text|The unique identifier for the resource This identifier is defined by the server|
|kind|text|Type of the resource Always compute#diskType for disk types|
|name|text|Name of the resource|
|region|text|URL of the region where the disk type resides Only applicable for regional resources You must specify this field as part of the HTTP request URL It is not settable as a field in the request body|
|self_link|text|Server-defined URL for the resource|
|valid_disk_size|text|An optional textual description of the valid disk size, such as "10GB-10TB"|
|zone|text|URL of the zone where the disk type resides You must specify this field as part of the HTTP request URL It is not settable as a field in the request body|
