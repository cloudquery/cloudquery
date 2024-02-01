# Table: aws_kms_key_policies

This table shows data for AWS Key Management Service (AWS KMS) Key Policies.

https://docs.aws.amazon.com/kms/latest/APIReference/API_GetKeyPolicy.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**key_arn**, **name**).
## Relations

This table depends on [aws_kms_keys](aws_kms_keys.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|key_arn|`utf8`|
|name|`utf8`|
|policy|`json`|