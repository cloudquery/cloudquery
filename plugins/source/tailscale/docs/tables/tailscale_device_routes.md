# Table: tailscale_device_routes

https://github.com/tailscale/tailscale/blob/main/api.md#device-routes-get

The primary key for this table is **_cq_id**.

## Relations

This table depends on [tailscale_devices](tailscale_devices.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|tailnet|String|
|device_id|String|
|advertised_routes|StringArray|
|enabled_routes|StringArray|