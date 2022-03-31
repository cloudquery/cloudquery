
# Table: aws_workspaces_workspaces
Describes a WorkSpace.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) for the workspaces workspace|
|bundle_id|text|The identifier of the bundle used to create the WorkSpace.|
|computer_name|text|The name of the WorkSpace, as seen by the operating system|
|directory_id|text|The identifier of the Directory Service directory for the WorkSpace.|
|error_code|text|The error code that is returned if the WorkSpace cannot be created.|
|error_message|text|The text of the error message that is returned if the WorkSpace cannot be created.|
|ip_address|text|The IP address of the WorkSpace.|
|modification_states|jsonb|The modification states of the WorkSpace.|
|root_volume_encryption_enabled|boolean|Indicates whether the data stored on the root volume is encrypted.|
|state|text|The operational state of the WorkSpace|
|subnet_id|text|The identifier of the subnet for the WorkSpace.|
|user_name|text|The user for the WorkSpace.|
|user_volume_encryption_enabled|boolean|Indicates whether the data stored on the user volume is encrypted.|
|volume_encryption_key|text|The symmetric KMS key used to encrypt data stored on your WorkSpace|
|id|text|The identifier of the WorkSpace.|
|compute_type_name|text|The compute type|
|root_volume_size_gib|integer|The size of the root volume|
|running_mode|text|The running mode|
|running_mode_auto_stop_timeout_in_minutes|integer|The time after a user logs off when WorkSpaces are automatically stopped. Configured in 60-minute intervals.|
|user_volume_size_gib|integer|The size of the user storage|
