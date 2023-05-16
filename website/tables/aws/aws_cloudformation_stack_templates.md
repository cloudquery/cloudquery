# Table: aws_cloudformation_stack_templates

This table shows data for AWS CloudFormation Stack Templates.

https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_GetTemplate.html

The primary key for this table is **stack_arn**.

## Relations

This table depends on [aws_cloudformation_stacks](aws_cloudformation_stacks).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|stack_arn (PK)|String|
|template_body|JSON|
|stages_available|StringArray|