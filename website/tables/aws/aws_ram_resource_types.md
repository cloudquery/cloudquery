# Table: aws_ram_resource_types

This table shows data for RAM Resource Types.

https://docs.aws.amazon.com/ram/latest/APIReference/API_ServiceNameAndResourceType.html

The composite primary key for this table is (**account_id**, **region**, **resource_type**, **service_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|resource_region_scope|`utf8`|
|resource_type (PK)|`utf8`|
|service_name (PK)|`utf8`|