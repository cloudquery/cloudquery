# Table: github_repository_sboms

This table shows data for Github Repository Sboms.

The composite primary key for this table is (**org**, **repository_id**).

## Relations

This table depends on [github_repositories](github_repositories.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|repository_id (PK)|`int64`|
|spdxid|`utf8`|
|spdx_version|`utf8`|
|creation_info|`json`|
|name|`utf8`|
|data_license|`utf8`|
|document_describes|`list<item: utf8, nullable>`|
|document_namespace|`utf8`|
|packages|`json`|