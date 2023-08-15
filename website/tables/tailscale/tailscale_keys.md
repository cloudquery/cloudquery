# Table: tailscale_keys

This table shows data for Tailscale Keys.

https://github.com/tailscale/tailscale/blob/main/api.md#keys

The composite primary key for this table is (**tailnet**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|tailnet (PK)|`utf8`|
|id (PK)|`utf8`|
|key|`utf8`|
|created|`timestamp[us, tz=UTC]`|
|expires|`timestamp[us, tz=UTC]`|
|capabilities|`json`|