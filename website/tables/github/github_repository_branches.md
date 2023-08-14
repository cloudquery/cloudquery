# Table: github_repository_branches

This table shows data for Github Repository Branches.

The composite primary key for this table is (**org**, **repository_id**, **name**).

## Relations

This table depends on [github_repositories](github_repositories).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|repository_id (PK)|`int64`|
|protection|`json`|
|name (PK)|`utf8`|
|commit|`json`|
|protected|`bool`|