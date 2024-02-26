# Table: okta_users

This table shows data for Okta Users.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|activated|`timestamp[us, tz=UTC]`|
|created|`timestamp[us, tz=UTC]`|
|credentials|`json`|
|id (PK)|`utf8`|
|last_login|`timestamp[us, tz=UTC]`|
|last_updated|`timestamp[us, tz=UTC]`|
|password_changed|`timestamp[us, tz=UTC]`|
|profile|`json`|
|status|`utf8`|
|status_changed|`timestamp[us, tz=UTC]`|
|transitioning_to_status|`utf8`|
|type|`json`|
|_embedded|`json`|
|_links|`json`|
|additional_properties|`json`|