# Table: gcp_apikeys_keys

This table shows data for GCP API Keys.

https://cloud.google.com/api-keys/docs/reference/rest/v2/projects.locations.keys#Key

The composite primary key for this table is (**project_id**, **uid**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name|`utf8`|
|uid (PK)|`utf8`|
|display_name|`utf8`|
|key_string|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|delete_time|`timestamp[us, tz=UTC]`|
|annotations|`json`|
|restrictions|`json`|
|etag|`utf8`|