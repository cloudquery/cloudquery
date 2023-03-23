# Table: aws_guardduty_detector_publishing_destinations

This table shows data for Amazon GuardDuty Detector Publishing Destinations.

https://docs.aws.amazon.com/guardduty/latest/APIReference/API_Member.html

The composite primary key for this table is (**detector_arn**, **destination_id**).

## Relations

This table depends on [aws_guardduty_detectors](aws_guardduty_detectors).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|detector_arn (PK)|String|
|destination_id (PK)|String|
|destination_type|String|
|status|String|