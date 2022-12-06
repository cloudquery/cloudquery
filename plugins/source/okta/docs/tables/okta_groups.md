# Table: okta_groups



The primary key for this table is **id**.

## Relations

The following tables depend on okta_groups:
  - [okta_group_users](okta_group_users.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|created|Timestamp|
|id (PK)|String|
|last_membership_updated|Timestamp|
|last_updated|Timestamp|
|object_class|StringArray|
|profile|JSON|
|type|String|