# Table: aws_rds_reserved_instances

This table shows data for Amazon Relational Database Service (RDS) Reserved Instances.

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_ReservedDBInstance.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|currency_code|`utf8`|
|db_instance_class|`utf8`|
|db_instance_count|`int64`|
|duration|`int64`|
|fixed_price|`float64`|
|lease_id|`utf8`|
|multi_az|`bool`|
|offering_type|`utf8`|
|product_description|`utf8`|
|recurring_charges|`json`|
|reserved_db_instance_arn|`utf8`|
|reserved_db_instance_id|`utf8`|
|reserved_db_instances_offering_id|`utf8`|
|start_time|`timestamp[us, tz=UTC]`|
|state|`utf8`|
|usage_price|`float64`|