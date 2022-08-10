
# Table: gcp_compute_projects
Represents a Project resource which is used to organize resources in a Google Cloud Platform environment
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|common_instance_metadata_fingerprint|text|Specifies a fingerprint for this resource|
|common_instance_metadata_items|jsonb|Array of key/value pairs The total size of all keys and values must be less than 512 KB|
|common_instance_metadata_kind|text|Type of the resource Always compute#metadata for metadata|
|creation_timestamp|timestamp without time zone|Creation timestamp in RFC3339 text format|
|default_network_tier|text|This signifies the default network tier used for configuring resources of the project and can only take the following values: PREMIUM, STANDARD Initially the default network tier is PREMIUM|
|default_service_account|text|Default service account used by VMs running in this project|
|description|text|An optional textual description of the resource|
|enabled_features|text[]|Restricted features enabled for use on this project|
|compute_project_id|text|The unique identifier for the resource This identifier is defined by the server This is not the project ID, and is just a unique ID used by Compute Engine to identify resources|
|kind|text|Type of the resource Always compute#project for projects|
|name|text|The project ID For example: my-example-project|
|self_link|text|Server-defined URL for the resource|
|usage_export_location_bucket_name|text|The name of an existing bucket in Cloud Storage where the usage report object is stored|
|usage_export_location_report_name_prefix|text|An optional prefix for the name of the usage report object stored in bucketName|
|xpn_project_status|text|The role this project has in a shared VPC configuration Currently, only projects with the host role, which is specified by the value HOST, are differentiated|
