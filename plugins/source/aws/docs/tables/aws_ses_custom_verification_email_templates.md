# Table: aws_ses_custom_verification_email_templates

https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetCustomVerificationEmailTemplate.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|failure_redirection_url|String|
|from_email_address|String|
|success_redirection_url|String|
|content|String|
|name|String|
|subject|String|