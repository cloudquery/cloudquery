# Table: vercel_project_envs

This table shows data for Vercel Project Envs.

The composite primary key for this table is (**project_id**, **id**).
It supports incremental syncs.
## Relations

This table depends on [vercel_projects](vercel_projects).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|id (PK)|`utf8`|
|target|`utf8`|
|type|`utf8`|
|key|`utf8`|
|value|`utf8`|
|configuration_id|`utf8`|
|git_branch|`utf8`|
|edge_config_id|`utf8`|
|edge_config_token_id|`utf8`|
|decrypted|`bool`|
|system|`bool`|
|created_at|`timestamp[us, tz=UTC]`|
|created_by|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|
|updated_by|`utf8`|