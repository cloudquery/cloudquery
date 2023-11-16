# Table: aws_stepfunctions_activities

This table shows data for Stepfunctions Activities.

https://docs.aws.amazon.com/step-functions/latest/apireference/API_ListActivities.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|activity_arn|`utf8`|
|creation_date|`timestamp[us, tz=UTC]`|
|name|`utf8`|