# Table: aws_quicksight_analyses

This table shows data for QuickSight Analyses.

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Analysis.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|analysis_id|`utf8`|
|arn|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|data_set_arns|`list<item: utf8, nullable>`|
|errors|`json`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|sheets|`json`|
|status|`utf8`|
|theme_arn|`utf8`|