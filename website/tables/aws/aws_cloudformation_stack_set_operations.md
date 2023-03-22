# Table: aws_cloudformation_stack_set_operations

This table shows data for AWS CloudFormation Stack Set Operations.

https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSetOperation.html

The composite primary key for this table is (**creation_timestamp**, **operation_id**, **stack_set_id**).

## Relations

This table depends on [aws_cloudformation_stack_sets](aws_cloudformation_stack_sets).

The following tables depend on aws_cloudformation_stack_set_operations:
  - [aws_cloudformation_stack_set_operation_results](aws_cloudformation_stack_set_operation_results)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|id|String|
|stack_set_arn|String|
|action|String|
|administration_role_arn|String|
|creation_timestamp (PK)|Timestamp|
|deployment_targets|JSON|
|end_timestamp|Timestamp|
|execution_role_name|String|
|operation_id (PK)|String|
|operation_preferences|JSON|
|retain_stacks|Bool|
|stack_set_drift_detection_details|JSON|
|stack_set_id (PK)|String|
|status|String|
|status_details|JSON|
|status_reason|String|