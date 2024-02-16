# Table: aws_elastictranscoder_presets

This table shows data for Amazon Elastic Transcoder Presets.

https://docs.aws.amazon.com/elastictranscoder/latest/developerguide/list-presets.html

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
|audio|`json`|
|container|`utf8`|
|description|`utf8`|
|id|`utf8`|
|name|`utf8`|
|thumbnails|`json`|
|type|`utf8`|
|video|`json`|