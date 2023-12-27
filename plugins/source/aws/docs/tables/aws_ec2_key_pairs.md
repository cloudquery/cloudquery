# Table: aws_ec2_key_pairs

This table shows data for Amazon Elastic Compute Cloud (EC2) Key Pairs.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_KeyPairInfo.html

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
|create_time|`timestamp[us, tz=UTC]`|
|key_fingerprint|`utf8`|
|key_name|`utf8`|
|key_pair_id|`utf8`|
|key_type|`utf8`|
|public_key|`utf8`|