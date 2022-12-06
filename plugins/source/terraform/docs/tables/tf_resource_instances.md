# Table: tf_resource_instances

Terraform resource instances

The primary key for this table is **_cq_id**.

## Relations
This table depends on [tf_resources](tf_resources.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|resource_name|String|
|instance_id|String|
|schema_version|Int|
|attributes|JSON|
|dependencies|StringArray|
|create_before_destroy|Bool|