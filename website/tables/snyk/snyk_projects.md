# Table: snyk_projects

This table shows data for Snyk Projects.

https://snyk.docs.apiary.io/#reference/projects/all-projects/list-all-projects

The composite primary key for this table is (**organization_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|organization_id (PK)|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|origin|`utf8`|