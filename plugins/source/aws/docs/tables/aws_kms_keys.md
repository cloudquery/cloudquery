# Table: aws_kms_keys

https://docs.aws.amazon.com/kms/latest/APIReference/API_KeyMetadata.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_kms_keys:
  - [aws_kms_key_grants](aws_kms_key_grants.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|rotation_enabled|Bool|
|tags|JSON|
|arn (PK)|String|
|replica_keys|JSON|
|key_id|String|
|aws_account_id|String|
|cloud_hsm_cluster_id|String|
|creation_date|Timestamp|
|custom_key_store_id|String|
|customer_master_key_spec|String|
|deletion_date|Timestamp|
|description|String|
|enabled|Bool|
|encryption_algorithms|StringArray|
|expiration_model|String|
|key_manager|String|
|key_spec|String|
|key_state|String|
|key_usage|String|
|mac_algorithms|StringArray|
|multi_region|Bool|
|multi_region_configuration|JSON|
|origin|String|
|pending_deletion_window_in_days|Int|
|signing_algorithms|StringArray|
|valid_to|Timestamp|