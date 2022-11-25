# Table: tf_data

Terraform meta data

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on tf_data:
  - [tf_resources](tf_resources.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|backend_type|String|
|backend_name|String|
|version|Int|
|terraform_version|String|
|serial|Int|
|lineage|String|