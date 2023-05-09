# Table: aws_cloudformation_template_summaries

This table shows data for AWS CloudFormation Template Summaries.

https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_GetTemplateSummary.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_cloudformation_stacks](aws_cloudformation_stacks).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|stack_id|String|
|metadata|JSON|
|capabilities|StringArray|
|capabilities_reason|String|
|declared_transforms|StringArray|
|description|String|
|parameters|JSON|
|resource_identifier_summaries|JSON|
|resource_types|StringArray|
|version|String|