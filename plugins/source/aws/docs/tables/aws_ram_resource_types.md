# Table: aws_ram_resource_types

https://docs.aws.amazon.com/ram/latest/APIReference/API_ServiceNameAndResourceType.html

The composite primary key for this table is (**account_id**, **resource_type**, **service_name**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|resource_region_scope|String|
|resource_type (PK)|String|
|service_name (PK)|String|