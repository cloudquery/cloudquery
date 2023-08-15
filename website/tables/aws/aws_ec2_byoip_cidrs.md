# Table: aws_ec2_byoip_cidrs

This table shows data for Amazon Elastic Compute Cloud (EC2) Bring your own IP addresses (BYOIP) CIDRs.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ByoipCidr.html

The composite primary key for this table is (**account_id**, **region**, **cidr**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|cidr (PK)|`utf8`|
|description|`utf8`|
|state|`utf8`|
|status_message|`utf8`|