# Table: aws_kms_key_policies

This table shows data for AWS Key Management Service (AWS KMS) Key Policies.

https://docs.aws.amazon.com/kms/latest/APIReference/API_GetKeyPolicy.html

The composite primary key for this table is (**key_arn**, **name**).

## Relations

This table depends on [aws_kms_keys](aws_kms_keys).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|key_arn (PK)|`utf8`|
|name (PK)|`utf8`|
|policy|`json`|