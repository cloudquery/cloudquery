# Table: aws_codepipeline_webhooks

https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_ListWebhookItem.html

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
|tags|JSON|
|definition|JSON|
|url|String|
|error_code|String|
|error_message|String|
|last_triggered|Timestamp|