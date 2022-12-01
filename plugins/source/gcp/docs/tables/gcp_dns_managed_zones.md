# Table: gcp_dns_managed_zones



The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|id (PK)|Int|
|cloud_logging_config|JSON|
|creation_time|String|
|description|String|
|dns_name|String|
|dnssec_config|JSON|
|forwarding_config|JSON|
|kind|String|
|labels|JSON|
|name|String|
|name_server_set|String|
|name_servers|StringArray|
|peering_config|JSON|
|private_visibility_config|JSON|
|reverse_lookup_config|JSON|
|service_directory_config|JSON|
|visibility|String|