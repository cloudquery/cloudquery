# Table: snyk_group_members

This table shows data for Snyk Group Members.

https://snyk.docs.apiary.io/#reference/groups/group-settings/list-all-members-in-a-group

The composite primary key for this table is (**group_id**, **id**).

## Relations

This table depends on [snyk_groups](snyk_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|group_id (PK)|String|
|id (PK)|String|
|name|String|
|username|String|
|email|String|
|orgs|JSON|
|group_role|String|