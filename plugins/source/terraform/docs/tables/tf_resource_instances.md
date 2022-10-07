# Table: tf_resource_instances

Terraform resource instances

The primary key for this table is **_cq_id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|resource_name|UUID|
|instance_id|String|
|schema_version|Int|
|attributes|JSON|
|dependencies|StringArray|
|create_before_destroy|Bool|