# Table: tf_resource_instances

This table shows data for Tf Resource Instances.

Terraform resource instances

The primary key for this table is **_cq_id**.

## Relations

This table depends on [tf_resources](tf_resources).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|resource_name|`utf8`|
|instance_id|`utf8`|
|schema_version|`int64`|
|attributes|`json`|
|dependencies|`list<item: utf8, nullable>`|
|create_before_destroy|`bool`|