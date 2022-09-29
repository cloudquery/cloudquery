# Table: tf_resources
Terraform resources

The primary key for this table is **_cq_id**.

## Relations
This table depends on [`tf_data`](tf_data.md).
The following tables depend on `tf_resources`:
  - [`tf_resource_instances`](tf_resource_instances.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|tf_data_cq_id|UUID|
|running_id|UUID|
|module|String|
|mode|String|
|type|String|
|name|String|
|provider_path|String|
|provider|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|