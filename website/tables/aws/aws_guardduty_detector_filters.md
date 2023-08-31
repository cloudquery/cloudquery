# Table: aws_guardduty_detector_filters

This table shows data for Amazon GuardDuty Detector Filters.

https://docs.aws.amazon.com/guardduty/latest/APIReference/API_GetFilter.html

The composite primary key for this table is (**detector_arn**, **name**).

## Relations

This table depends on [aws_guardduty_detectors](aws_guardduty_detectors).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|detector_arn (PK)|`utf8`|
|action|`utf8`|
|finding_criteria|`json`|
|name (PK)|`utf8`|
|description|`utf8`|
|rank|`int64`|
|tags|`json`|