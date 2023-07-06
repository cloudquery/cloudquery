# Table: aws_guardduty_detector_members

This table shows data for Amazon GuardDuty Detector Members.

https://docs.aws.amazon.com/guardduty/latest/APIReference/API_Member.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_guardduty_detectors](aws_guardduty_detectors).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|region|`utf8`|
|detector_arn|`utf8`|
|account_id|`utf8`|
|email|`utf8`|
|master_id|`utf8`|
|relationship_status|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|
|administrator_id|`utf8`|
|detector_id|`utf8`|
|invited_at|`timestamp[us, tz=UTC]`|