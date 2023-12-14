# Table: aws_ec2_instance_credit_specifications

This table shows data for Amazon Elastic Compute Cloud (EC2) Instance Credit Specifications.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceCreditSpecification.html

The composite primary key for this table is (**account_id**, **region**, **instance_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|cpu_credits|`utf8`|
|instance_id (PK)|`utf8`|