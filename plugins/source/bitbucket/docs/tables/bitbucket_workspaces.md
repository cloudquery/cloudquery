# Table: bitbucket_workspaces

https://developer.atlassian.com/cloud/bitbucket/rest/api-group-workspaces/#api-group-workspaces

The primary key for this table is **uuid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|uuid (PK)|`uuid`|
|name|`utf8`|
|type|`utf8`|
|slug|`utf8`|
|is_private|`bool`|
|created_on|`timestamp[ms, tz=Z]`|
|updated_on|`timestamp[ms, tz=Z]`|
|links|`json`|