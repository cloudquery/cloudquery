# Table: aws_codepipeline_pipelines

This table shows data for Codepipeline Pipelines.

https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_GetPipeline.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|metadata|`json`|
|pipeline|`json`|