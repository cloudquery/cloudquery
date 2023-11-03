# Table: aws_cloudformation_stack_instance_summaries

This table shows data for AWS CloudFormation Stack Instance Summaries.

https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackInstanceSummary.html

The composite primary key for this table is (**stack_set_arn**, **stack_set_id**).

## Relations

This table depends on [aws_cloudformation_stack_sets](aws_cloudformation_stack_sets.md).

The following tables depend on aws_cloudformation_stack_instance_summaries:
  - [aws_cloudformation_stack_instance_resource_drifts](aws_cloudformation_stack_instance_resource_drifts.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|id|`utf8`|
|stack_set_arn (PK)|`utf8`|
|account|`utf8`|
|drift_status|`utf8`|
|last_drift_check_timestamp|`timestamp[us, tz=UTC]`|
|last_operation_id|`utf8`|
|organizational_unit_id|`utf8`|
|stack_id|`utf8`|
|stack_instance_status|`json`|
|stack_set_id (PK)|`utf8`|
|status|`utf8`|
|status_reason|`utf8`|