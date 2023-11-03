# Table: aws_ses_templates

This table shows data for Amazon Simple Email Service (SES) Templates.

https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetEmailTemplate.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|template_name|`utf8`|
|html|`utf8`|
|subject|`utf8`|
|text|`utf8`|
|created_timestamp|`timestamp[us, tz=UTC]`|