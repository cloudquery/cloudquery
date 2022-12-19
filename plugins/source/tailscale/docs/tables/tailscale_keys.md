# Table: tailscale_keys

https://pkg.go.dev/github.com/tailscale/tailscale-client-go/tailscale#Key

The primary key for this table is **tailnet**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|tailnet (PK)|String|
|id|String|
|key|String|
|created|Timestamp|
|expires|Timestamp|
|capabilities|JSON|