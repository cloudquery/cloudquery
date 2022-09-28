# Table: tf_resource_instances
Terraform resource instances

The primary key for this table is **_cq_id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|tf_resource_cq_id|UUID|
|resource_id|UUID|
|instance_id|String|
|schema_version|Int|
|attributes|JSON|
|dependencies|StringArray|
|create_before_destroy|Bool|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|