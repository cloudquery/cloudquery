# Table: render_service_deploys

This table shows data for Render Service Deploys.

The primary key for this table is **id**.

## Relations

This table depends on [render_services](render_services.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|service_id|`utf8`|
|id (PK)|`utf8`|
|commit|`json`|
|status|`utf8`|
|finished_at|`timestamp[us, tz=UTC]`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|