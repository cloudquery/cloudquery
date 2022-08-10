
# Table: aws_lambda_function_concurrency_configs
Details about the provisioned concurrency configuration for a function alias or version.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_cq_id|uuid|Unique CloudQuery ID of aws_lambda_functions table (FK)|
|allocated_provisioned_concurrent_executions|integer|The amount of provisioned concurrency allocated.|
|available_provisioned_concurrent_executions|integer|The amount of provisioned concurrency available.|
|function_arn|text|The Amazon Resource Name (ARN) of the alias or version.|
|last_modified|timestamp without time zone|The date and time that a user last updated the configuration, in ISO 8601 format (https://www.iso.org/iso-8601-date-and-time-format.html).|
|requested_provisioned_concurrent_executions|integer|The amount of provisioned concurrency requested.|
|status|text|The status of the allocation process.|
|status_reason|text|For failed allocations, the reason that provisioned concurrency could not be allocated.|
