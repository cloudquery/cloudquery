# Table: gcp_dns_policies


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|project_id|String|
|id (PK)|Int|
|alternative_name_server_config|JSON|
|description|String|
|enable_inbound_forwarding|Bool|
|enable_logging|Bool|
|kind|String|
|name|String|
|networks|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|