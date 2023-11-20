# Table: aws_cloudformation_stack_set_operation_results

This table shows data for AWS CloudFormation Stack Set Operation Results.

https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSetOperationResultSummary.html

The 'request_account_id' and 'request_region' columns are added to show the account and region of where the request was made from.

The composite primary key for this table is (**request_account_id**, **request_region**, **stack_set_arn**, **operation_id**, **account**, **region**).

## Relations

This table depends on [aws_cloudformation_stack_set_operations](aws_cloudformation_stack_set_operations.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id (PK)|`utf8`|
|request_region (PK)|`utf8`|
|stack_set_arn (PK)|`utf8`|
|operation_id (PK)|`utf8`|
|account (PK)|`utf8`|
|account_gate_result|`json`|
|organizational_unit_id|`utf8`|
|region (PK)|`utf8`|
|status|`utf8`|
|status_reason|`utf8`|