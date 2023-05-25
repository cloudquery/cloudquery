# Table: okta_users

This table shows data for Okta Users.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|activated|timestamp[us, tz=UTC]|
|created|timestamp[us, tz=UTC]|
|credentials|extension_type<storage=binary>|
|id (PK)|utf8|
|last_login|timestamp[us, tz=UTC]|
|last_updated|timestamp[us, tz=UTC]|
|password_changed|timestamp[us, tz=UTC]|
|profile|extension_type<storage=binary>|
|status|utf8|
|status_changed|timestamp[us, tz=UTC]|
|transitioning_to_status|utf8|
|type|extension_type<storage=binary>|
|_embedded|extension_type<storage=binary>|
|_links|extension_type<storage=binary>|
|additional_properties|extension_type<storage=binary>|