# Table: aws_cloudformation_stack_instance_resource_drifts

This table shows data for AWS CloudFormation Stack Instance Resource Drifts.

https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSetOperationResultSummary.html.
The 'request_account_id' and 'request_region' columns are added to show the account and region of where the request was made from.

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_cloudformation_stack_instance_summaries](aws_cloudformation_stack_instance_summaries).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|request_region|`utf8`|
|stack_set_arn|`utf8`|
|logical_resource_id|`utf8`|
|resource_type|`utf8`|
|stack_id|`utf8`|
|stack_resource_drift_status|`utf8`|
|timestamp|`timestamp[us, tz=UTC]`|
|physical_resource_id|`utf8`|
|physical_resource_id_context|`json`|
|property_differences|`json`|