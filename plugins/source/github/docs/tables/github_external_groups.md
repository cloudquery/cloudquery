# Table: github_external_groups

This table shows data for Github External Groups.

The composite primary key for this table is (**org**, **group_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|group_id (PK)|`int64`|
|group_name|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|
|teams|`json`|
|members|`json`|