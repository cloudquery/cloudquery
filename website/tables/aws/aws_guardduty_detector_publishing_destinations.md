# Table: aws_guardduty_detector_publishing_destinations

This table shows data for Amazon GuardDuty Detector Publishing Destinations.

https://docs.aws.amazon.com/guardduty/latest/APIReference/API_Member.html

The composite primary key for this table is (**detector_arn**, **destination_id**).

## Relations

This table depends on [aws_guardduty_detectors](aws_guardduty_detectors).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|detector_arn (PK)|`utf8`|
|destination_id (PK)|`utf8`|
|destination_type|`utf8`|
|status|`utf8`|