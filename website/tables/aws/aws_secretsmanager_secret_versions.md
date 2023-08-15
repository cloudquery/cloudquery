# Table: aws_secretsmanager_secret_versions

This table shows data for AWS Secrets Manager Secret Versions.

https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_ListSecretVersionIds.html

The composite primary key for this table is (**secret_arn**, **version_id**).

## Relations

This table depends on [aws_secretsmanager_secrets](aws_secretsmanager_secrets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|secret_arn (PK)|`utf8`|
|created_date|`timestamp[us, tz=UTC]`|
|kms_key_ids|`list<item: utf8, nullable>`|
|last_accessed_date|`timestamp[us, tz=UTC]`|
|version_id (PK)|`utf8`|
|version_stages|`list<item: utf8, nullable>`|