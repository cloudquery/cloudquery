# Table: tf_resources

Terraform resources

The primary key for this table is **_cq_id**.

## Relations
This table depends on [tf_data](tf_data.md).

The following tables depend on tf_resources:
  - [tf_resource_instances](tf_resource_instances.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|data_backend_name|String|
|module|String|
|mode|String|
|type|String|
|name|String|
|provider_path|String|
|provider|String|