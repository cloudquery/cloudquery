# Table: aws_codepipeline_pipelines

This table shows data for Codepipeline Pipelines.

https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_GetPipeline.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|metadata|`json`|
|pipeline|`json`|
|result_metadata|`json`|