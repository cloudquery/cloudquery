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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|id|String|
|arn (PK)|String|
|tags|JSON|
|administration_role_arn|String|
|auto_deployment|JSON|
|capabilities|StringArray|
|description|String|
|execution_role_name|String|
|managed_execution|JSON|
|organizational_unit_ids|StringArray|
|parameters|JSON|
|permission_model|String|
|stack_set_arn|String|
|stack_set_drift_detection_details|JSON|
|stack_set_id|String|
|stack_set_name|String|
|status|String|
|template_body|String|