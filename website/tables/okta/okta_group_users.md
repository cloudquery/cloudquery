# Table: okta_group_users

This table shows data for Okta Group Users.

The composite primary key for this table is (**group_id**, **id**).

## Relations

This table depends on [okta_groups](okta_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|group_id (PK)|`utf8`|
|id (PK)|`utf8`|