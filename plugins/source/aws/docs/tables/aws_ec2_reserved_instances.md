# Table: aws_ec2_reserved_instances

This table shows data for Amazon Elastic Compute Cloud (EC2) Reserved Instances.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ReservedInstances.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|availability_zone|`utf8`|
|currency_code|`utf8`|
|duration|`int64`|
|end|`timestamp[us, tz=UTC]`|
|fixed_price|`float64`|
|instance_count|`int64`|
|instance_tenancy|`utf8`|
|instance_type|`utf8`|
|offering_class|`utf8`|
|offering_type|`utf8`|
|product_description|`utf8`|
|recurring_charges|`json`|
|reserved_instances_id|`utf8`|
|scope|`utf8`|
|start|`timestamp[us, tz=UTC]`|
|state|`utf8`|
|usage_price|`float64`|