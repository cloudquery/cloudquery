# Table: aws_secretsmanager_secrets

https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_ListSecrets.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_secretsmanager_secrets:
  - [aws_secretsmanager_secret_versions](aws_secretsmanager_secret_versions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|policy|JSON|
|tags|JSON|
|created_date|Timestamp|
|deleted_date|Timestamp|
|description|String|
|kms_key_id|String|
|last_accessed_date|Timestamp|
|last_changed_date|Timestamp|
|last_rotated_date|Timestamp|
|name|String|
|next_rotation_date|Timestamp|
|owning_service|String|
|primary_region|String|
|replication_status|JSON|
|rotation_enabled|Bool|
|rotation_lambda_arn|String|
|rotation_rules|JSON|
|version_ids_to_stages|JSON|