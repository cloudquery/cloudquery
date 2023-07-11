# Table: tf_data

This table shows data for Tf Data.

Terraform meta data

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on tf_data:
  - [tf_resources](tf_resources)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|backend_type|`utf8`|
|backend_name|`utf8`|
|version|`int64`|
|terraform_version|`utf8`|
|serial|`int64`|
|lineage|`utf8`|