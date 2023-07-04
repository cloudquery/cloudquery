# Table: gcp_binaryauthorization_assertors

This table shows data for GCP Binary Authorization Assertors.

https://cloud.google.com/binary-authorization/docs/reference/rest/v1/projects.attestors#Attestor

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|description|`utf8`|
|update_time|`timestamp[us, tz=UTC]`|