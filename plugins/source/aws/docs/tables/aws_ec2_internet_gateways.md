# Table: aws_ec2_internet_gateways

This table shows data for Amazon Elastic Compute Cloud (EC2) Internet Gateways.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InternetGateway.html

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
|attachments|`json`|
|internet_gateway_id|`utf8`|
|owner_id|`utf8`|