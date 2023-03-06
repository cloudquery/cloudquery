# Table: aws_quicksight_analyses

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Analysis.html

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
|analysis_id|String|
|arn (PK)|String|
|created_time|Timestamp|
|data_set_arns|StringArray|
|errors|JSON|
|last_updated_time|Timestamp|
|name|String|
|sheets|JSON|
|status|String|
|theme_arn|String|