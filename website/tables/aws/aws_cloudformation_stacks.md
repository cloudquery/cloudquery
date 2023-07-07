# Table: aws_cloudformation_stacks

This table shows data for AWS CloudFormation Stacks.

https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_Stack.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_cloudformation_stacks:
  - [aws_cloudformation_stack_resources](aws_cloudformation_stack_resources)
  - [aws_cloudformation_stack_templates](aws_cloudformation_stack_templates)
  - [aws_cloudformation_template_summaries](aws_cloudformation_template_summaries)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|id|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|creation_time|`timestamp[us, tz=UTC]`|
|stack_name|`utf8`|
|stack_status|`utf8`|
|capabilities|`list<item: utf8, nullable>`|
|change_set_id|`utf8`|
|deletion_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|disable_rollback|`bool`|
|drift_information|`json`|
|enable_termination_protection|`bool`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|notification_arns|`list<item: utf8, nullable>`|
|outputs|`json`|
|parameters|`json`|
|parent_id|`utf8`|
|role_arn|`utf8`|
|rollback_configuration|`json`|
|root_id|`utf8`|
|stack_id|`utf8`|
|stack_status_reason|`utf8`|
|timeout_in_minutes|`int64`|