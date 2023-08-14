# Table: aws_glue_ml_transforms

This table shows data for Glue ML Transforms.

https://docs.aws.amazon.com/glue/latest/webapi/API_MLTransform.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_glue_ml_transforms:
  - [aws_glue_ml_transform_task_runs](aws_glue_ml_transform_task_runs)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|schema|`json`|
|created_on|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|evaluation_metrics|`json`|
|glue_version|`utf8`|
|input_record_tables|`json`|
|label_count|`int64`|
|last_modified_on|`timestamp[us, tz=UTC]`|
|max_capacity|`float64`|
|max_retries|`int64`|
|name|`utf8`|
|number_of_workers|`int64`|
|parameters|`json`|
|role|`utf8`|
|status|`utf8`|
|timeout|`int64`|
|transform_encryption|`json`|
|transform_id|`utf8`|
|worker_type|`utf8`|