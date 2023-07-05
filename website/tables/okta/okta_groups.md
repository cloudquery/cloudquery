# Table: okta_groups

This table shows data for Okta Groups.

The primary key for this table is **id**.

## Relations

The following tables depend on okta_groups:
  - [okta_group_users](okta_group_users)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|created|`timestamp[us, tz=UTC]`|
|id (PK)|`utf8`|
|last_membership_updated|`timestamp[us, tz=UTC]`|
|last_updated|`timestamp[us, tz=UTC]`|
|object_class|`list<item: utf8, nullable>`|
|profile|`json`|
|type|`utf8`|
|_embedded|`json`|
|_links|`json`|
|additional_properties|`json`|