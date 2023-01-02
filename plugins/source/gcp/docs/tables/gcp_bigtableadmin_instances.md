# Table: gcp_bigtableadmin_instances

https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances#Instance

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|display_name|String|
|instance_state|Int|
|instance_type|Int|
|labels|JSON|