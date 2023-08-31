# Table: aws_cloudformation_stack_set_operations

This table shows data for AWS CloudFormation Stack Set Operations.

https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSetOperation.html

The composite primary key for this table is (**stack_set_arn**, **creation_timestamp**, **operation_id**).

## Relations

This table depends on [aws_cloudformation_stack_sets](aws_cloudformation_stack_sets).

The following tables depend on aws_cloudformation_stack_set_operations:
  - [aws_cloudformation_stack_set_operation_results](aws_cloudformation_stack_set_operation_results)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|id|`utf8`|
|stack_set_arn (PK)|`utf8`|
|action|`utf8`|
|administration_role_arn|`utf8`|
|creation_timestamp (PK)|`timestamp[us, tz=UTC]`|
|deployment_targets|`json`|
|end_timestamp|`timestamp[us, tz=UTC]`|
|execution_role_name|`utf8`|
|operation_id (PK)|`utf8`|
|operation_preferences|`json`|
|retain_stacks|`bool`|
|stack_set_drift_detection_details|`json`|
|stack_set_id|`utf8`|
|status|`utf8`|
|status_details|`json`|
|status_reason|`utf8`|