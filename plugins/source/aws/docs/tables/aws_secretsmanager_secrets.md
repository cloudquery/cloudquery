
# Table: aws_secretsmanager_secrets
A structure that contains the details about a secret
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|policy|jsonb|A JSON-formatted string that describes the permissions that are associated with the attached secret.|
|replication_status|jsonb|A replication object consisting of a RegionReplicationStatus object and includes a Region, KMSKeyId, status, and status message.|
|arn|text|The Amazon Resource Name (ARN) of the secret|
|created_date|timestamp without time zone|The date and time when a secret was created.|
|deleted_date|timestamp without time zone|The date and time the deletion of the secret occurred|
|description|text|The user-provided description of the secret.|
|kms_key_id|text|The ARN or alias of the Amazon Web Services KMS customer master key (CMK) used to encrypt the SecretString and SecretBinary fields in each version of the secret|
|last_accessed_date|timestamp without time zone|The last date that this secret was accessed|
|last_changed_date|timestamp without time zone|The last date and time that this secret was modified in any way.|
|last_rotated_date|timestamp without time zone|The most recent date and time that the Secrets Manager rotation process was successfully completed|
|name|text|The friendly name of the secret|
|owning_service|text|Returns the name of the service that created the secret.|
|primary_region|text|The Region where Secrets Manager originated the secret.|
|rotation_enabled|boolean|Indicates whether automatic, scheduled rotation is enabled for this secret.|
|rotation_lambda_arn|text|The ARN of an Amazon Web Services Lambda function invoked by Secrets Manager to rotate and expire the secret either automatically per the schedule or manually by a call to RotateSecret.|
|rotation_rules_automatically_after_days|bigint|Specifies the number of days between automatic scheduled rotations of the secret|
|secret_versions_to_stages|jsonb|A list of all of the currently assigned SecretVersionStage staging labels and the SecretVersionId attached to each one|
|tags|jsonb|The list of user-defined tags associated with the secret|
