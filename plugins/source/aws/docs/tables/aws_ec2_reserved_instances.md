# Table: aws_ec2_reserved_instances

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ReservedInstances.html

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
|availability_zone|String|
|currency_code|String|
|duration|Int|
|end|Timestamp|
|fixed_price|Float|
|instance_count|Int|
|instance_tenancy|String|
|instance_type|String|
|offering_class|String|
|offering_type|String|
|product_description|String|
|recurring_charges|JSON|
|reserved_instances_id|String|
|scope|String|
|start|Timestamp|
|state|String|
|tags|JSON|
|usage_price|Float|