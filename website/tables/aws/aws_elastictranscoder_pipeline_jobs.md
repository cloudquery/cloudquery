# Table: aws_elastictranscoder_pipeline_jobs

This table shows data for Amazon Elastic Transcoder Pipeline Jobs.

https://docs.aws.amazon.com/elastictranscoder/latest/developerguide/list-jobs-by-pipeline.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_elastictranscoder_pipelines](aws_elastictranscoder_pipelines).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|id|`utf8`|
|input|`json`|
|inputs|`json`|
|output|`json`|
|output_key_prefix|`utf8`|
|outputs|`json`|
|pipeline_id|`utf8`|
|playlists|`json`|
|status|`utf8`|
|timing|`json`|
|user_metadata|`json`|