# Table: aws_ec2_egress_only_internet_gateways

This table shows data for Amazon Elastic Compute Cloud (EC2) Egress Only Internet Gateways.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_EgressOnlyInternetGateway.html

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
|attachments|`json`|
|egress_only_internet_gateway_id|`utf8`|