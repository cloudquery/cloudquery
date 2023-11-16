# Table: gcp_iam_deny_policies

This table shows data for GCP IAM Deny Policies.

https://cloud.google.com/iam/docs/reference/rest/v2beta/policies#Policy

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|uid|`utf8`|
|kind|`utf8`|
|display_name|`utf8`|
|annotations|`json`|
|etag|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|delete_time|`timestamp[us, tz=UTC]`|
|rules|`json`|
|managing_authority|`utf8`|