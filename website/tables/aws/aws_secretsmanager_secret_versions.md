# Table: aws_secretsmanager_secret_versions

This table shows data for AWS Secretsmanager Secret Versions.

https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_ListSecretVersionIds.html

The composite primary key for this table is (**secret_arn**, **version_id**).

## Relations

This table depends on [aws_secretsmanager_secrets](aws_secretsmanager_secrets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|secret_arn (PK)|String|
|created_date|Timestamp|
|kms_key_ids|StringArray|
|last_accessed_date|Timestamp|
|version_id (PK)|String|
|version_stages|StringArray|