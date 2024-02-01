# Table: aws_kms_aliases

This table shows data for AWS Key Management Service (AWS KMS) Aliases.

https://docs.aws.amazon.com/kms/latest/APIReference/API_AliasListEntry.html

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
|alias_arn|`utf8`|
|alias_name|`utf8`|
|creation_date|`timestamp[us, tz=UTC]`|
|last_updated_date|`timestamp[us, tz=UTC]`|
|target_key_id|`utf8`|