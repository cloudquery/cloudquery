# Table: render_services

This table shows data for Render Services.

The primary key for this table is **id**.

## Relations

The following tables depend on render_services:
  - [render_service_custom_domains](render_service_custom_domains.md)
  - [render_service_deploys](render_service_deploys.md)
  - [render_service_jobs](render_service_jobs.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|id (PK)|`utf8`|
|auto_deploy|`bool`|
|branch|`utf8`|
|build_filter|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|image|`json`|
|name|`utf8`|
|notify_on_fail|`utf8`|
|owner_id|`utf8`|
|repo|`utf8`|
|root_dir|`utf8`|
|slug|`utf8`|
|suspended|`utf8`|
|suspenders|`list<item: utf8, nullable>`|
|type|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|
|service_details|`json`|