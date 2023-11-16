# Table: gcp_resourcemanager_project_tag_values

This table shows data for GCP Resourcemanager Project Tag Values.

https://cloud.google.com/resource-manager/reference/rest/v3/tagValues/list

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_resourcemanager_project_tag_keys](gcp_resourcemanager_project_tag_keys.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|parent|`utf8`|
|short_name|`utf8`|
|namespaced_name|`utf8`|
|description|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|etag|`utf8`|