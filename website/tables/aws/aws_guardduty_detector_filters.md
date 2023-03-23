# Table: aws_guardduty_detector_filters

This table shows data for Amazon GuardDuty Detector Filters.

https://docs.aws.amazon.com/guardduty/latest/APIReference/API_GetFilter.html

The primary key for this table is **name**.

## Relations

This table depends on [aws_guardduty_detectors](aws_guardduty_detectors).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|detector_arn|String|
|action|String|
|finding_criteria|JSON|
|name (PK)|String|
|description|String|
|rank|Int|
|tags|JSON|
|result_metadata|JSON|