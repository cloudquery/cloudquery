# Table: okta_application_group_assignments

This table shows data for Okta Application Group Assignments.

The composite primary key for this table is (**app_id**, **id**).

## Relations

This table depends on [okta_applications](okta_applications).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|app_id (PK)|utf8|
|id (PK)|utf8|
|last_updated|timestamp[us, tz=UTC]|
|priority|int64|
|profile|extension_type<storage=binary>|
|_embedded|extension_type<storage=binary>|
|_links|extension_type<storage=binary>|
|additional_properties|extension_type<storage=binary>|