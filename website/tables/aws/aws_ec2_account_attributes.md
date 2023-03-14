# Table: aws_ec2_account_attributes

This table shows data for Amazon Elastic Compute Cloud (EC2) Account Attributes.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AccountAttribute.html

The composite primary key for this table is (**account_id**, **attribute_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|partition|String|
|attribute_name (PK)|String|
|attribute_values|JSON|