# Table: tailscale_device_routes

https://pkg.go.dev/github.com/tailscale/tailscale-client-go/tailscale#DeviceRoutes

The composite primary key for this table is (**tailnet**, **device_id**).

## Relations

This table depends on [tailscale_devices](tailscale_devices.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|tailnet (PK)|String|
|device_id (PK)|String|
|advertised|StringArray|
|enabled|StringArray|