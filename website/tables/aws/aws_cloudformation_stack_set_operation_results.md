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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|request_account_id|String|
|request_region|String|
|operation_id|String|
|stack_set_arn|String|
|account|String|
|account_gate_result|JSON|
|organizational_unit_id|String|
|region|String|
|status|String|
|status_reason|String|