# Table: aws_ec2_byoip_cidrs

This table shows data for Amazon Elastic Compute Cloud (EC2) Bring your own IP addresses (BYOIP) CIDRs.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ByoipCidr.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **cidr**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|cidr|`utf8`|
|asn_associations|`json`|
|description|`utf8`|
|state|`utf8`|
|status_message|`utf8`|