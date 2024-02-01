# Table: aws_ram_resource_types

This table shows data for RAM Resource Types.

https://docs.aws.amazon.com/ram/latest/APIReference/API_ServiceNameAndResourceType.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **resource_type**, **service_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|resource_region_scope|`utf8`|
|resource_type|`utf8`|
|service_name|`utf8`|