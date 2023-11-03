# Table: aws_ses_custom_verification_email_templates

This table shows data for Amazon Simple Email Service (SES) Custom Verification Email Templates.

https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetCustomVerificationEmailTemplate.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|failure_redirection_url|`utf8`|
|from_email_address|`utf8`|
|success_redirection_url|`utf8`|
|content|`utf8`|
|name|`utf8`|
|subject|`utf8`|