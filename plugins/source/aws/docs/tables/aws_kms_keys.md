# Table: aws_kms_keys

This table shows data for AWS Key Management Service (AWS KMS) Keys.

https://docs.aws.amazon.com/kms/latest/APIReference/API_KeyMetadata.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_kms_keys:
  - [aws_kms_key_grants](aws_kms_key_grants.md)
  - [aws_kms_key_policies](aws_kms_key_policies.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|rotation_enabled|`bool`|
|tags|`json`|
|arn|`utf8`|
|replica_keys|`json`|
|key_id|`utf8`|
|aws_account_id|`utf8`|
|cloud_hsm_cluster_id|`utf8`|
|creation_date|`timestamp[us, tz=UTC]`|
|custom_key_store_id|`utf8`|
|customer_master_key_spec|`utf8`|
|deletion_date|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|enabled|`bool`|
|encryption_algorithms|`list<item: utf8, nullable>`|
|expiration_model|`utf8`|
|key_manager|`utf8`|
|key_spec|`utf8`|
|key_state|`utf8`|
|key_usage|`utf8`|
|mac_algorithms|`list<item: utf8, nullable>`|
|multi_region|`bool`|
|multi_region_configuration|`json`|
|origin|`utf8`|
|pending_deletion_window_in_days|`int64`|
|signing_algorithms|`list<item: utf8, nullable>`|
|valid_to|`timestamp[us, tz=UTC]`|
|xks_key_configuration|`json`|