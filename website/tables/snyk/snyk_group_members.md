# Table: snyk_group_members

This table shows data for Snyk Group Members.

https://snyk.docs.apiary.io/#reference/groups/group-settings/list-all-members-in-a-group

The composite primary key for this table is (**group_id**, **id**).

## Relations

This table depends on [snyk_groups](snyk_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|group_id (PK)|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|username|`utf8`|
|email|`utf8`|
|orgs|`json`|
|group_role|`utf8`|