# Table: aws_cloudformation_stack_sets

This table shows data for AWS CloudFormation Stack Sets.

https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSet.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_cloudformation_stack_sets:
  - [aws_cloudformation_stack_set_operations](aws_cloudformation_stack_set_operations)

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
|administration_role_arn|`utf8`|
|auto_deployment|`json`|
|capabilities|`list<item: utf8, nullable>`|
|description|`utf8`|
|execution_role_name|`utf8`|
|managed_execution|`json`|
|organizational_unit_ids|`list<item: utf8, nullable>`|
|parameters|`json`|
|permission_model|`utf8`|
|regions|`list<item: utf8, nullable>`|
|stack_set_arn|`utf8`|
|stack_set_drift_detection_details|`json`|
|stack_set_id|`utf8`|
|stack_set_name|`utf8`|
|status|`utf8`|
|template_body|`utf8`|