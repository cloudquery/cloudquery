# Table: tailscale_dns_nameservers

This table shows data for Tailscale DNS Nameservers.

https://github.com/tailscale/tailscale/blob/main/api.md#tailnet-dns-preferences-get

The composite primary key for this table is (**tailnet**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|tailnet (PK)|`utf8`|
|name (PK)|`utf8`|