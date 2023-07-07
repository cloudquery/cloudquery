# Table: tf_resources

This table shows data for Tf Resources.

Terraform resources

The primary key for this table is **_cq_id**.

## Relations

This table depends on [tf_data](tf_data).

The following tables depend on tf_resources:
  - [tf_resource_instances](tf_resource_instances)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|data_backend_name|`utf8`|
|module|`utf8`|
|mode|`utf8`|
|type|`utf8`|
|name|`utf8`|
|provider_path|`utf8`|
|provider|`utf8`|