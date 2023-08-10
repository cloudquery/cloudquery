# Table: digitalocean_apps

This table shows data for DigitalOcean Apps.

https://docs.digitalocean.com/reference/api/api-reference/#operation/apps_get

The primary key for this table is **id**.

## Relations

The following tables depend on digitalocean_apps:
  - [digitalocean_apps_alerts](digitalocean_apps_alerts)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`uuid`|
|owner_uuid|`utf8`|
|spec|`json`|
|last_deployment_active_at|`timestamp[us, tz=UTC]`|
|default_ingress|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|active_deployment|`json`|
|in_progress_deployment|`json`|
|pending_deployment|`json`|
|last_deployment_created_at|`timestamp[us, tz=UTC]`|
|live_url|`utf8`|
|region|`json`|
|tier_slug|`utf8`|
|live_url_base|`utf8`|
|live_domain|`utf8`|
|domains|`json`|
|pinned_deployment|`json`|
|build_config|`json`|
|project_id|`utf8`|