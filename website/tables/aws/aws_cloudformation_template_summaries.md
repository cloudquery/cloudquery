# Table: aws_cloudformation_template_summaries

This table shows data for AWS CloudFormation Template Summaries.

https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_GetTemplateSummary.html

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
|stack_id|String|
|stack_arn (PK)|String|
|metadata|JSON|
|capabilities|StringArray|
|capabilities_reason|String|
|declared_transforms|StringArray|
|description|String|
|parameters|JSON|
|resource_identifier_summaries|JSON|
|resource_types|StringArray|
|version|String|