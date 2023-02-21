# Table: tailscale_devices

https://github.com/tailscale/tailscale/blob/main/api.md#tailnet-devices-get

The composite primary key for this table is (**tailnet**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|tailnet (PK)|String|
|advertised_routes|StringArray|
|enabled_routes|StringArray|
|addresses|StringArray|
|name|String|
|id (PK)|String|
|authorized|Bool|
|user|String|
|tags|StringArray|
|key_expiry_disabled|Bool|
|blocks_incoming_connections|Bool|
|client_version|String|
|created|Timestamp|
|expires|Timestamp|
|hostname|String|
|is_external|Bool|
|last_seen|Timestamp|
|machine_key|String|
|node_key|String|
|os|String|
|update_available|Bool|