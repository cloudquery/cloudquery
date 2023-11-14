# Table: bitbucket_repositories

https://developer.atlassian.com/cloud/bitbucket/rest/api-group-repositories/#api-group-repositories

The primary key for this table is **uuid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|uuid (PK)|`uuid`|
|type|`utf8`|
|full_name|`utf8`|
|is_private|`bool`|
|parent|`json`|
|scm|`utf8`|
|name|`utf8`|
|description|`utf8`|
|created_on|`timestamp[ms, tz=Z]`|
|updated_on|`timestamp[ms, tz=Z]`|
|size|`int64`|
|language|`utf8`|
|has_issues|`bool`|
|has_wiki|`bool`|
|fork_policy|`utf8`|
|project|`json`|
|main_branch|`json`|
|links|`json`|