# Table: aws_cloudformation_stack_set_operation_results

This table shows data for AWS CloudFormation Stack Set Operation Results.

https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSetOperationResultSummary.html.
The 'request_account_id' and 'request_region' columns are added to show the account and region of where the request was made from.

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_cloudformation_stack_set_operations](aws_cloudformation_stack_set_operations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|request_region|`utf8`|
|operation_id|`utf8`|
|stack_set_arn|`utf8`|
|account|`utf8`|
|account_gate_result|`json`|
|organizational_unit_id|`utf8`|
|region|`utf8`|
|status|`utf8`|
|status_reason|`utf8`|