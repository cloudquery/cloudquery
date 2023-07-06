# Table: okta_application_users

This table shows data for Okta Application Users.

The composite primary key for this table is (**app_id**, **id**).

## Relations

This table depends on [okta_applications](okta_applications).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|app_id (PK)|`utf8`|
|created|`timestamp[us, tz=UTC]`|
|credentials|`json`|
|external_id|`utf8`|
|id (PK)|`utf8`|
|last_sync|`timestamp[us, tz=UTC]`|
|last_updated|`timestamp[us, tz=UTC]`|
|password_changed|`timestamp[us, tz=UTC]`|
|profile|`json`|
|scope|`utf8`|
|status|`utf8`|
|status_changed|`timestamp[us, tz=UTC]`|
|sync_state|`utf8`|
|_embedded|`json`|
|_links|`json`|
|additional_properties|`json`|