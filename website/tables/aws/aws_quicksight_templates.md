# Table: aws_quicksight_templates

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_TemplateSummary.html

The composite primary key for this table is (**account_id**, **region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|tags|JSON|
|arn (PK)|String|
|created_time|Timestamp|
|last_updated_time|Timestamp|
|latest_version_number|Int|
|name|String|
|template_id|String|