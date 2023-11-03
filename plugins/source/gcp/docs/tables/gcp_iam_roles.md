# Table: gcp_iam_roles

This table shows data for GCP IAM Roles.

https://cloud.google.com/iam/docs/reference/rest/v1/roles#Role

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|title|`utf8`|
|description|`utf8`|
|included_permissions|`list<item: utf8, nullable>`|
|stage|`utf8`|
|etag|`binary`|
|deleted|`bool`|