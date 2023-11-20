# Table: aws_cloudformation_stack_resources

This table shows data for AWS CloudFormation Stack Resources.

https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackResourceSummary.html

The composite primary key for this table is (**stack_id**, **logical_resource_id**).

## Relations

This table depends on [aws_cloudformation_stacks](aws_cloudformation_stacks.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|stack_id (PK)|`utf8`|
|last_updated_timestamp|`timestamp[us, tz=UTC]`|
|logical_resource_id (PK)|`utf8`|
|resource_status|`utf8`|
|resource_type|`utf8`|
|drift_information|`json`|
|module_info|`json`|
|physical_resource_id|`utf8`|
|resource_status_reason|`utf8`|