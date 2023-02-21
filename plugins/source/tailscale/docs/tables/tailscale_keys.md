# Table: tailscale_keys

https://github.com/tailscale/tailscale/blob/main/api.md#keys

The composite primary key for this table is (**tailnet**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|tailnet (PK)|String|
|id (PK)|String|
|key|String|
|created|Timestamp|
|expires|Timestamp|
|capabilities|JSON|