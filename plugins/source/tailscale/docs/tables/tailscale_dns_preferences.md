# Table: tailscale_dns_preferences

https://github.com/tailscale/tailscale/blob/main/api.md#tailnet-dns-preferences-get

The primary key for this table is **tailnet**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|tailnet (PK)|String|
|magic_dns|Bool|