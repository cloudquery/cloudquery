# Table: aws_guardduty_detector_members

This table shows data for Amazon GuardDuty Detector Members.

https://docs.aws.amazon.com/guardduty/latest/APIReference/API_Member.html

The composite primary key for this table is (**detector_arn**, **account_id**).

## Relations

This table depends on [aws_guardduty_detectors](aws_guardduty_detectors.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|request_region|`utf8`|
|detector_arn (PK)|`utf8`|
|account_id (PK)|`utf8`|
|email|`utf8`|
|master_id|`utf8`|
|relationship_status|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|
|administrator_id|`utf8`|
|detector_id|`utf8`|
|invited_at|`timestamp[us, tz=UTC]`|