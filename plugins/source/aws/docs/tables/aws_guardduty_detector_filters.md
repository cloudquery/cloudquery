# Table: aws_guardduty_detector_filters

This table shows data for Amazon GuardDuty Detector Filters.

https://docs.aws.amazon.com/guardduty/latest/APIReference/API_GetFilter.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**request_account_id**, **request_region**, **detector_arn**, **arn**, **name**).
## Relations

This table depends on [aws_guardduty_detectors](aws_guardduty_detectors.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|request_region|`utf8`|
|detector_arn|`utf8`|
|arn|`utf8`|
|action|`utf8`|
|finding_criteria|`json`|
|name|`utf8`|
|description|`utf8`|
|rank|`int64`|
|tags|`json`|