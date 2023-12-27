# Table: aws_cloudformation_stack_instance_resource_drifts

This table shows data for AWS CloudFormation Stack Instance Resource Drifts.

https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackInstanceResourceDriftsSummary.html.
The 'request_account_id' and 'request_region' columns are added to show the account and region of where the request was made from.

The composite primary key for this table is (**stack_set_arn**, **operation_id**, **logical_resource_id**, **stack_id**, **physical_resource_id**).

## Relations

This table depends on [aws_cloudformation_stack_instance_summaries](aws_cloudformation_stack_instance_summaries.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|request_region|`utf8`|
|stack_set_arn (PK)|`utf8`|
|operation_id (PK)|`utf8`|
|logical_resource_id (PK)|`utf8`|
|resource_type|`utf8`|
|stack_id (PK)|`utf8`|
|stack_resource_drift_status|`utf8`|
|timestamp|`timestamp[us, tz=UTC]`|
|physical_resource_id (PK)|`utf8`|
|physical_resource_id_context|`json`|
|property_differences|`json`|