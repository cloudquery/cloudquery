
# Table: aws_lambda_functions
AWS Lambda is a serverless compute service that lets you run code without provisioning or managing servers, creating workload-aware cluster scaling logic, maintaining event integrations, or managing runtimes
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|policy_document|jsonb|The resource-based policy.|
|policy_revision_id|text|A unique identifier for the current revision of the policy.|
|code_signing_allowed_publishers_version_arns|text[]|The Amazon Resource Name (ARN) for each of the signing profiles. A signing profile defines a trusted user who can sign a code package.|
|code_signing_config_arn|text|The Amazon Resource Name (ARN) of the Code signing configuration.|
|code_signing_config_id|text|Unique identifier for the Code signing configuration.|
|code_signing_policies_untrusted_artifact_on_deployment|text|Code signing configuration policy for deployment validation failure.|
|code_signing_description|text|Code signing configuration description.|
|code_signing_last_modified|timestamp without time zone|The date and time that the Code signing configuration was last modified, in ISO-8601 format (YYYY-MM-DDThh:mm:ss.sTZD).|
|code_image_uri|text|URI of a container image in the Amazon ECR registry.|
|code_location|text|A presigned URL that you can use to download the deployment package.|
|code_repository_type|text|The service that's hosting the file.|
|code_resolved_image_uri|text|The resolved URI for the image.|
|concurrency_reserved_concurrent_executions|integer|The number of concurrent executions that are reserved for this function|
|architectures|text[]|The instruction set architecture that the function supports|
|code_sha256|text|The SHA256 hash of the function's deployment package.|
|code_size|bigint|The size of the function's deployment package, in bytes.|
|dead_letter_config_target_arn|text|The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.|
|description|text|The function's description.|
|environment_error_code|text|The error code.|
|environment_error_message|text|The error message.|
|environment_variables|jsonb|Environment variable key-value pairs.|
|ephemeral_storage_size|integer|The size of the function’s /tmp directory.|
|arn|text|The function's Amazon Resource Name (ARN).|
|name|text|The name of the function.|
|handler|text|The function that Lambda calls to begin executing your function.|
|error_code|text|Error code.|
|error_message|text|Error message.|
|image_config_command|text[]|Specifies parameters that you want to pass in with ENTRYPOINT.|
|image_config_entry_point|text[]|Specifies the entry point to their application, which is typically the location of the runtime executable.|
|image_config_working_directory|text|Specifies the working directory.|
|kms_key_arn|text|The KMS key that's used to encrypt the function's environment variables|
|last_modified|timestamp without time zone|The date and time that the function was last updated, in ISO-8601 format (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).|
|last_update_status|text|The status of the last update that was performed on the function|
|last_update_status_reason|text|The reason for the last update that was performed on the function.|
|last_update_status_reason_code|text|The reason code for the last update that was performed on the function.|
|master_arn|text|For Lambda@Edge functions, the ARN of the main function.|
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
|tags|jsonb|The function's tags (https://docs.aws.amazon.com/lambda/latest/dg/tagging.html).|
