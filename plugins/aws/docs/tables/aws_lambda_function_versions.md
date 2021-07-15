
# Table: aws_lambda_function_versions
Details about a function's configuration. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_cq_id|uuid|Unique CloudQuery ID of aws_lambda_functions table (FK)|
|code_sha256|text|The SHA256 hash of the function's deployment package.|
|code_size|bigint|The size of the function's deployment package, in bytes.|
|dead_letter_config_target_arn|text|The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.|
|description|text|The function's description.|
|environment_error_error_code|text|The error code.|
|environment_error_message|text|The error message.|
|environment_variables|jsonb|Environment variable key-value pairs.|
|function_arn|text|The function's Amazon Resource Name (ARN).|
|function_name|text|The name of the function.|
|handler|text|The function that Lambda calls to begin executing your function.|
|error_code|text|Error code.|
|error_message|text|Error message.|
|image_config_command|text[]|Specifies parameters that you want to pass in with ENTRYPOINT.|
|image_config_entry_point|text[]|Specifies the entry point to their application, which is typically the location of the runtime executable.|
|image_config_working_directory|text|Specifies the working directory.|
|kms_key_arn|text|The KMS key that's used to encrypt the function's environment variables|
|last_modified|text|The date and time that the function was last updated, in ISO-8601 format (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).|
|last_update_status|text|The status of the last update that was performed on the function|
|last_update_status_reason|text|The reason for the last update that was performed on the function.|
|last_update_status_reason_code|text|The reason code for the last update that was performed on the function.|
|master_arn|text|For Lambda@Edge functions, the ARN of the master function.|
|memory_size|integer|The amount of memory available to the function at runtime.|
|package_type|text|The type of deployment package|
|revision_id|text|The latest updated revision of the function or alias.|
|role|text|The function's execution role.|
|runtime|text|The runtime environment for the Lambda function.|
|signing_job_arn|text|The ARN of the signing job.|
|signing_profile_version_arn|text|The ARN of the signing profile version.|
|state|text|The current state of the function|
|state_reason|text|The reason for the function's current state.|
|state_reason_code|text|The reason code for the function's current state|
|timeout|integer|The amount of time in seconds that Lambda allows a function to run before stopping it.|
|tracing_config_mode|text|The tracing mode.|
|version|text|The version of the Lambda function.|
|vpc_config_security_group_ids|text[]|A list of VPC security groups IDs.|
|vpc_config_subnet_ids|text[]|A list of VPC subnet IDs.|
|vpc_config_vpc_id|text|The ID of the VPC.|
