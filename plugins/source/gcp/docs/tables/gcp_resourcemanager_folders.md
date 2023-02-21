# Table: gcp_resourcemanager_folders

https://cloud.google.com/resource-manager/reference/rest/v3/folders#Folder

The composite primary key for this table is (**organization_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|organization_id (PK)|String|
|name (PK)|String|
|parent|String|
|display_name|String|
|state|String|
|create_time|Timestamp|
|update_time|Timestamp|
|delete_time|Timestamp|
|etag|String|