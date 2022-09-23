# Table: aws_codepipeline_webhooks


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|definition|JSON|
|url|String|
|error_code|String|
|error_message|String|
|last_triggered|Timestamp|
|tags|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|