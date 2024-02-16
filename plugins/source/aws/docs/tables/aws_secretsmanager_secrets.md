# Table: aws_secretsmanager_secrets

This table shows data for AWS Secrets Manager Secrets.

https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_ListSecrets.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_secretsmanager_secrets:
  - [aws_secretsmanager_secret_versions](aws_secretsmanager_secret_versions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|policy|`json`|
|tags|`json`|
|created_date|`timestamp[us, tz=UTC]`|
|deleted_date|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|kms_key_id|`utf8`|
|last_accessed_date|`timestamp[us, tz=UTC]`|
|last_changed_date|`timestamp[us, tz=UTC]`|
|last_rotated_date|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|next_rotation_date|`timestamp[us, tz=UTC]`|
|owning_service|`utf8`|
|primary_region|`utf8`|
|replication_status|`json`|
|rotation_enabled|`bool`|
|rotation_lambda_arn|`utf8`|
|rotation_rules|`json`|
|version_ids_to_stages|`json`|