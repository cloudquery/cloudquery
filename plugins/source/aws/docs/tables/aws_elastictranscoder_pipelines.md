# Table: aws_elastictranscoder_pipelines

This table shows data for Amazon Elastic Transcoder Pipelines.

https://docs.aws.amazon.com/elastictranscoder/latest/developerguide/list-pipelines.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_elastictranscoder_pipelines:
  - [aws_elastictranscoder_pipeline_jobs](aws_elastictranscoder_pipeline_jobs.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|aws_kms_key_arn|`utf8`|
|content_config|`json`|
|id|`utf8`|
|input_bucket|`utf8`|
|name|`utf8`|
|notifications|`json`|
|output_bucket|`utf8`|
|role|`utf8`|
|status|`utf8`|
|thumbnail_config|`json`|