# Table: aws_cloudformation_stack_templates

This table shows data for AWS CloudFormation Stack Templates.

https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_GetTemplate.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **stack_arn**.
## Relations

This table depends on [aws_cloudformation_stacks](aws_cloudformation_stacks.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|stack_arn|`utf8`|
|template_body|`json`|
|template_body_text|`utf8`|
|stages_available|`list<item: utf8, nullable>`|