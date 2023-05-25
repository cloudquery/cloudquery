# Table: gcp_dns_policies

This table shows data for GCP DNS Policies.

https://cloud.google.com/dns/docs/reference/v1/policies#resource

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|project_id|utf8|
|alternative_name_server_config|json|
|description|utf8|
|enable_inbound_forwarding|bool|
|enable_logging|bool|
|id (PK)|int64|
|kind|utf8|
|name|utf8|
|networks|json|