# Table: aws_kms_key_policies

https://docs.aws.amazon.com/kms/latest/APIReference/API_GetKeyPolicy.html

The composite primary key for this table is (**key_arn**, **name**).

## Relations

This table depends on [aws_kms_keys](aws_kms_keys.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|key_arn (PK)|String|
|name (PK)|String|
|policy|JSON|