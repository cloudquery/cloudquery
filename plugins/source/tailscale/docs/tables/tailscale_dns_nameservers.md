# Table: tailscale_dns_nameservers

https://github.com/tailscale/tailscale/blob/main/api.md#tailnet-dns-preferences-get

The composite primary key for this table is (**tailnet**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|tailnet (PK)|String|
|name (PK)|String|