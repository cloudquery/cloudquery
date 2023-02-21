# Table: gcp_dns_policies

https://cloud.google.com/dns/docs/reference/v1/policies#resource

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|alternative_name_server_config|JSON|
|description|String|
|enable_inbound_forwarding|Bool|
|enable_logging|Bool|
|id (PK)|Int|
|kind|String|
|name|String|
|networks|JSON|