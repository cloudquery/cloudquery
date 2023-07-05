# Table: tailscale_dns_preferences

This table shows data for Tailscale Domain Name System (DNS) Preferences.

https://github.com/tailscale/tailscale/blob/main/api.md#tailnet-dns-preferences-get

The primary key for this table is **tailnet**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|tailnet (PK)|`utf8`|
|magic_dns|`bool`|