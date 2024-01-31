# Table: aws_secretsmanager_secret_versions

This table shows data for AWS Secrets Manager Secret Versions.

https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_ListSecretVersionIds.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**secret_arn**, **version_id**).
## Relations

This table depends on [aws_secretsmanager_secrets](aws_secretsmanager_secrets.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|secret_arn|`utf8`|
|created_date|`timestamp[us, tz=UTC]`|
|kms_key_ids|`list<item: utf8, nullable>`|
|last_accessed_date|`timestamp[us, tz=UTC]`|
|version_id|`utf8`|
|version_stages|`list<item: utf8, nullable>`|