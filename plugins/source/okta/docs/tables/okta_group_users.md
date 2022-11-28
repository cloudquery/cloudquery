# Table: okta_group_users

The primary key for this table is **_cq_id**.

## Relations

This table depends on [okta_groups](okta_groups.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|group_id|String|
|id|String|