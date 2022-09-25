# Table: azure_network_watchers


The primary key for this table is **id**.

## Relations
The following tables depend on `azure_network_watchers`:
  - [`azure_network_flow_logs`](azure_network_flow_logs.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|etag|String|
|provisioning_state|String|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|