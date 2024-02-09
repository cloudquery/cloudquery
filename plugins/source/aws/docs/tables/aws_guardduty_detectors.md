# Table: aws_guardduty_detectors

This table shows data for Amazon GuardDuty Detectors.

https://docs.aws.amazon.com/guardduty/latest/APIReference/API_GetDetector.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**request_account_id**, **request_region**, **arn**, **id**).
## Relations

The following tables depend on aws_guardduty_detectors:
  - [aws_guardduty_detector_filters](aws_guardduty_detector_filters.md)
  - [aws_guardduty_detector_findings](aws_guardduty_detector_findings.md)
  - [aws_guardduty_detector_intel_sets](aws_guardduty_detector_intel_sets.md)
  - [aws_guardduty_detector_ip_sets](aws_guardduty_detector_ip_sets.md)
  - [aws_guardduty_detector_members](aws_guardduty_detector_members.md)
  - [aws_guardduty_detector_publishing_destinations](aws_guardduty_detector_publishing_destinations.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|request_region|`utf8`|
|arn|`utf8`|
|service_role|`utf8`|
|status|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|data_sources|`json`|
|features|`json`|
|finding_publishing_frequency|`utf8`|
|tags|`json`|
|updated_at|`timestamp[us, tz=UTC]`|
|id|`utf8`|