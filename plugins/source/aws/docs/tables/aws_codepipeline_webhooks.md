# Table: aws_codepipeline_webhooks

This table shows data for Codepipeline Webhooks.

https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_ListWebhookItem.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|definition|`json`|
|url|`utf8`|
|error_code|`utf8`|
|error_message|`utf8`|
|last_triggered|`timestamp[us, tz=UTC]`|