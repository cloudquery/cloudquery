# Table: aws_ec2_account_attributes

This table shows data for Amazon Elastic Compute Cloud (EC2) Account Attributes.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AccountAttribute.html

The composite primary key for this table is (**account_id**, **attribute_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|partition|`utf8`|
|attribute_name (PK)|`utf8`|
|attribute_values|`json`|