
# Table: aws_kms_keys
Contains information about each entry in the key list.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|rotation_enabled|boolean|specifies whether key rotation is enabled.|
|cloud_hsm_cluster_id|text|The cluster ID of the AWS CloudHSM cluster that contains the key material for the CMK|
|creation_date|timestamp without time zone|The date and time when the CMK was created.|
|custom_key_store_id|text|A unique identifier for the custom key store.|
|customer_master_key_spec|text|Describes the type of key material in the CMK.|
|deletion_date|timestamp without time zone|he date and time after which AWS KMS deletes the CMK. This value is present only when KeyState is PendingDeletion.|
|description|text|The description of the CMK.|
|enabled|boolean|Specifies whether the CMK is enabled.|
|encryption_algorithms|text[]|The encryption algorithms that the CMK supports.|
|expiration_model|text|Specifies whether the CMK's key material expires.|
|manager|text|The manager of the CMK.|
|key_state|text|The current status of the CMK.|
|key_usage|text|The cryptographic operations for which you can use the CMK.|
|origin|text|The source of the CMK's key material.|
|signing_algorithms|text[]|The signing algorithms that the CMK supports.|
|valid_to|timestamp without time zone|The time at which the imported key material expires.|
|arn|text|ARN of the key|
|key_id|text|Unique identifier of the key|
