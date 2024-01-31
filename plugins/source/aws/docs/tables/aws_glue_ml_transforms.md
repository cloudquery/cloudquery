# Table: aws_glue_ml_transforms

This table shows data for Glue ML Transforms.

https://docs.aws.amazon.com/glue/latest/webapi/API_MLTransform.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_glue_ml_transforms:
  - [aws_glue_ml_transform_task_runs](aws_glue_ml_transform_task_runs.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|schema|`json`|
|tags|`json`|
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