
# Table: aws_kms_keys
Contains metadata about a KMS key
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|rotation_enabled|boolean|Specifies whether key rotation is enabled.|
|tags|jsonb|Key tags.|
|id|text|The globally unique identifier for the KMS key.|
|aws_account_id|text|The twelve-digit account ID of the Amazon Web Services account that owns the KMS key.|
|arn|text|The Amazon Resource Name (ARN) of the KMS key|
|cloud_hsm_cluster_id|text|The cluster ID of the CloudHSM cluster that contains the key material for the KMS key|
|creation_date|timestamp without time zone|The date and time when the KMS key was created.|
|custom_key_store_id|text|A unique identifier for the custom key store that contains the KMS key|
|deletion_date|timestamp without time zone|The date and time after which KMS deletes this KMS key|
|description|text|The description of the KMS key.|
|enabled|boolean|Specifies whether the KMS key is enabled|
|encryption_algorithms|text[]|The encryption algorithms that the KMS key supports|
|expiration_model|text|Specifies whether the KMS key's key material expires|
|manager|text|The manager of the KMS key|
|key_spec|text|Describes the type of key material in the KMS key.|
|key_state|text|The current status of the KMS key|
|key_usage|text|The cryptographic operations for which you can use the KMS key.|
|mac_algorithms|text[]|The message authentication code (MAC) algorithm that the HMAC KMS key supports. This value is present only when the KeyUsage of the KMS key is GENERATE_VERIFY_MAC.|
|multi_region|boolean|Indicates whether the KMS key is a multi-Region (True) or regional (False) key. This value is True for multi-Region primary and replica keys and False for regional KMS keys|
|multi_region_key_type|text|Indicates whether the KMS key is a PRIMARY or REPLICA key.|
|primary_key_arn|text|Displays the key ARN of a primary or replica key of a multi-Region key.|
|primary_key_region|text|Displays the Amazon Web Services Region of a primary or replica key in a multi-Region key.|
|replica_keys|jsonb|displays the key ARNs and Regions of all replica keys|
|origin|text|The source of the key material for the KMS key|
|pending_deletion_window_in_days|integer|The waiting period before the primary key in a multi-Region key is deleted|
|signing_algorithms|text[]|The signing algorithms that the KMS key supports|
|valid_to|timestamp without time zone|The time at which the imported key material expires|
