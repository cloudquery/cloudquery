# Table: aws_guardduty_detector_ip_sets

This table shows data for Amazon GuardDuty Detector IP Sets.

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
|format|String|
|location|String|
|name (PK)|String|
|status|String|
|tags|JSON|