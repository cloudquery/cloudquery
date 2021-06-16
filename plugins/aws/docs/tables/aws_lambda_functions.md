
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
|code_image_uri|text||
|code_location|text||
|code_repository_type|text||
|code_resolved_image_uri|text||
|concurrency_reserved_concurrent_executions|integer||
|code_sha256|text||
|code_size|bigint||
|dead_letter_config_target_arn|text||
|description|text||
|environment_error_code|text||
|environment_error_message|text||
|environment_variables|jsonb||
|function_arn|text||
|function_name|text||
|handler|text||
|error_code|text||
|error_message|text||
|image_config_command|text[]||
|image_config_entry_point|text[]||
|image_config_working_directory|text||
|kms_key_arn|text||
|last_modified|text||
|last_update_status|text||
|last_update_status_reason|text||
|last_update_status_reason_code|text||
|master_arn|text||
|memory_size|integer||
|package_type|text||
|revision_id|text||
|role|text||
|runtime|text||
|signing_job_arn|text||
|signing_profile_version_arn|text||
|state|text||
|state_reason|text||
|state_reason_code|text||
|timeout|integer||
|tracing_config_mode|text||
|version|text||
|vpc_config_security_group_ids|text[]||
|vpc_config_subnet_ids|text[]||
|vpc_config_vpc_id|text||
|tags|jsonb|The function's tags (https://docs.aws.amazon.com/lambda/latest/dg/tagging.html).|
