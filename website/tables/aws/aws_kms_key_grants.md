# Table: aws_kms_key_grants

This table shows data for AWS Key Management Service (AWS KMS) Key Grants.

https://docs.aws.amazon.com/kms/latest/APIReference/API_GrantListEntry.html

The composite primary key for this table is (**key_arn**, **grant_id**).

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
|grant_id (PK)|`utf8`|
|constraints|`json`|
|creation_date|`timestamp[us, tz=UTC]`|
|grantee_principal|`utf8`|
|issuing_account|`utf8`|
|key_id|`utf8`|
|name|`utf8`|
|operations|`list<item: utf8, nullable>`|
|retiring_principal|`utf8`|