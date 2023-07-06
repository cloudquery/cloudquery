# Table: aws_cloudformation_template_summaries

This table shows data for AWS CloudFormation Template Summaries.

https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_GetTemplateSummary.html

The primary key for this table is **stack_arn**.

## Relations

This table depends on [aws_cloudformation_stacks](aws_cloudformation_stacks).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|stack_id|`utf8`|
|stack_arn (PK)|`utf8`|
|metadata|`json`|
|capabilities|`list<item: utf8, nullable>`|
|capabilities_reason|`utf8`|
|declared_transforms|`list<item: utf8, nullable>`|
|description|`utf8`|
|parameters|`json`|
|resource_identifier_summaries|`json`|
|resource_types|`list<item: utf8, nullable>`|
|version|`utf8`|