# Table: aws_kms_key_grants

https://docs.aws.amazon.com/kms/latest/APIReference/API_GrantListEntry.html

The composite primary key for this table is (**key_arn**, **grant_id**).

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
|grant_id (PK)|String|
|constraints|JSON|
|creation_date|Timestamp|
|grantee_principal|String|
|issuing_account|String|
|key_id|String|
|name|String|
|operations|StringArray|
|retiring_principal|String|