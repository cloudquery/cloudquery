# Table: aws_quicksight_analyses

This table shows data for QuickSight Analyses.

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Analysis.html

The composite primary key for this table is (**account_id**, **region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|tags|`json`|
|analysis_id|`utf8`|
|arn (PK)|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|data_set_arns|`list<item: utf8, nullable>`|
|errors|`json`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|sheets|`json`|
|status|`utf8`|
|theme_arn|`utf8`|