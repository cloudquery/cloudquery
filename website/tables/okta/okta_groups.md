# Table: okta_groups

This table shows data for Okta Groups.

The primary key for this table is **id**.

## Relations

The following tables depend on okta_groups:
  - [okta_group_users](okta_group_users)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|created|Timestamp|
|last_membership_updated|Timestamp|
|last_updated|Timestamp|
|object_class|StringArray|
|profile|JSON|
|type|String|
|_embedded|JSON|
|_links|JSON|
|additional_properties|JSON|