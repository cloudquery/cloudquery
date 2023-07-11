# Table: gcp_resourcemanager_folders

This table shows data for GCP Resourcemanager Folders.

https://cloud.google.com/resource-manager/reference/rest/v3/folders#Folder

The composite primary key for this table is (**organization_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|organization_id (PK)|`utf8`|
|name (PK)|`utf8`|
|parent|`utf8`|
|display_name|`utf8`|
|state|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|delete_time|`timestamp[us, tz=UTC]`|
|etag|`utf8`|