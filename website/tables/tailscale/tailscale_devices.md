# Table: tailscale_devices

This table shows data for Tailscale Devices.

https://github.com/tailscale/tailscale/blob/main/api.md#tailnet-devices-get

The composite primary key for this table is (**tailnet**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|tailnet (PK)|`utf8`|
|advertised_routes|`list<item: utf8, nullable>`|
|enabled_routes|`list<item: utf8, nullable>`|
|addresses|`list<item: utf8, nullable>`|
|name|`utf8`|
|id (PK)|`utf8`|
|authorized|`bool`|
|user|`utf8`|
|tags|`list<item: utf8, nullable>`|
|key_expiry_disabled|`bool`|
|blocks_incoming_connections|`bool`|
|client_version|`utf8`|
|created|`timestamp[us, tz=UTC]`|
|expires|`timestamp[us, tz=UTC]`|
|hostname|`utf8`|
|is_external|`bool`|
|last_seen|`timestamp[us, tz=UTC]`|
|machine_key|`utf8`|
|node_key|`utf8`|
|os|`utf8`|
|update_available|`bool`|