# Table: aws_rds_reserved_instances

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_ReservedDBInstance.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|currency_code|String|
|db_instance_class|String|
|db_instance_count|Int|
|duration|Int|
|fixed_price|Float|
|lease_id|String|
|multi_az|Bool|
|offering_type|String|
|product_description|String|
|recurring_charges|JSON|
|reserved_db_instance_arn|String|
|reserved_db_instance_id|String|
|reserved_db_instances_offering_id|String|
|start_time|Timestamp|
|state|String|
|usage_price|Float|