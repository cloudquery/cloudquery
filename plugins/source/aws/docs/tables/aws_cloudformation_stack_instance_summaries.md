# Table: aws_cloudformation_stack_instance_summaries

This table shows data for AWS CloudFormation Stack Instance Summaries.

https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackInstanceSummary.html

**Note**: Sometimes the stack instance ID may be unavailable in the API (i.e., the instance is in a bad state), so it will have value of `N/A`.

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**stack_set_arn**, **stack_id**, **stack_set_id**).
## Relations

This table depends on [aws_cloudformation_stack_sets](aws_cloudformation_stack_sets.md).

The following tables depend on aws_cloudformation_stack_instance_summaries:
  - [aws_cloudformation_stack_instance_resource_drifts](aws_cloudformation_stack_instance_resource_drifts.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|id|`utf8`|
|stack_set_arn|`utf8`|
|account|`utf8`|
|drift_status|`utf8`|
|last_drift_check_timestamp|`timestamp[us, tz=UTC]`|
|last_operation_id|`utf8`|
|organizational_unit_id|`utf8`|
|stack_id|`utf8`|
|stack_instance_status|`json`|
|stack_set_id|`utf8`|
|status|`utf8`|
|status_reason|`utf8`|