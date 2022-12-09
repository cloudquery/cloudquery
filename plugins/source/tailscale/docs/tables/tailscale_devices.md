# Table: tailscale_devices

https://pkg.go.dev/github.com/tailscale/tailscale-client-go/tailscale#Device

The primary key for this table is **tailnet**.

## Relations

The following tables depend on tailscale_devices:
  - [tailscale_device_routes](tailscale_device_routes.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|tailnet (PK)|String|
|addresses|StringArray|
|name|String|
|id|String|
|authorized|Bool|
|user|String|
|tags|StringArray|
|key_expiry_disabled|Bool|
|blocks_incoming_connections|Bool|
|client_version|String|
|created|JSON|
|expires|JSON|
|hostname|String|
|is_external|Bool|
|last_seen|JSON|
|machine_key|String|
|node_key|String|
|os|String|
|update_available|Bool|