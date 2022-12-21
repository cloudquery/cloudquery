# Table: aws_elastictranscoder_pipeline_jobs

https://docs.aws.amazon.com/elastictranscoder/latest/developerguide/list-jobs-by-pipeline.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_elastictranscoder_pipelines](aws_elastictranscoder_pipelines.md).

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
|id|String|
|input|JSON|
|inputs|JSON|
|output|JSON|
|output_key_prefix|String|
|outputs|JSON|
|pipeline_id|String|
|playlists|JSON|
|status|String|
|timing|JSON|
|user_metadata|JSON|