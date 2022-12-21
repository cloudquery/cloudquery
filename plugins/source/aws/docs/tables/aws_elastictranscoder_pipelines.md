# Table: aws_elastictranscoder_pipelines

https://docs.aws.amazon.com/elastictranscoder/latest/developerguide/list-pipelines.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_elastictranscoder_pipelines:
  - [aws_elastictranscoder_pipeline_jobs](aws_elastictranscoder_pipeline_jobs.md)

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
|aws_kms_key_arn|String|
|content_config|JSON|
|id|String|
|input_bucket|String|
|name|String|
|notifications|JSON|
|output_bucket|String|
|role|String|
|status|String|
|thumbnail_config|JSON|