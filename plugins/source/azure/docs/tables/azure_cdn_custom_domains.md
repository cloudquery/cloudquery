# Table: azure_cdn_custom_domains



The primary key for this table is **id**.

## Relations
This table depends on [azure_cdn_endpoints](azure_cdn_endpoints.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|cdn_endpoint_id|String|
|host_name|String|
|resource_state|String|
|custom_https_provisioning_state|String|
|custom_https_provisioning_substate|String|
|validation_data|String|
|provisioning_state|String|
|id (PK)|String|
|name|String|
|type|String|
|system_data|JSON|