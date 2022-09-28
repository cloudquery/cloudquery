# Table: tf_data
Terraform meta data

The primary key for this table is **_cq_id**.

## Relations
The following tables depend on `tf_data`:
  - [`tf_resources`](tf_resources.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|backend_type|String|
|backend_name|String|
|version|Int|
|terraform_version|String|
|serial|Int|
|lineage|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|