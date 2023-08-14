# Table: aws_cloudformation_stack_templates

This table shows data for AWS CloudFormation Stack Templates.

https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_GetTemplate.html

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
|stack_arn (PK)|`utf8`|
|template_body|`json`|
|template_body_text|`utf8`|
|stages_available|`list<item: utf8, nullable>`|