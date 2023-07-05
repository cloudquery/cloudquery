# Table: okta_application_group_assignments

This table shows data for Okta Application Group Assignments.

The composite primary key for this table is (**app_id**, **id**).

## Relations

This table depends on [okta_applications](okta_applications).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|app_id (PK)|`utf8`|
|id (PK)|`utf8`|
|last_updated|`timestamp[us, tz=UTC]`|
|priority|`int64`|
|profile|`json`|
|_embedded|`json`|
|_links|`json`|
|additional_properties|`json`|