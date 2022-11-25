# Table: okta_group_users



The composite primary key for this table is (**group_id**, **id**).

## Relations
This table depends on [okta_groups](okta_groups.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|group_id (PK)|String|
|id (PK)|String|