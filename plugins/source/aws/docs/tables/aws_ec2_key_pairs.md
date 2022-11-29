# Table: aws_ec2_key_pairs

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_KeyPairInfo.html

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
|create_time|Timestamp|
|key_fingerprint|String|
|key_name|String|
|key_pair_id|String|
|key_type|String|
|public_key|String|
|tags|JSON|