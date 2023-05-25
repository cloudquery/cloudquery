# Table: aws_guardduty_detectors

This table shows data for Amazon GuardDuty Detectors.

https://docs.aws.amazon.com/guardduty/latest/APIReference/API_GetDetector.html

The composite primary key for this table is (**account_id**, **region**, **id**).

## Relations

The following tables depend on aws_guardduty_detectors:
  - [aws_guardduty_detector_filters](aws_guardduty_detector_filters)
  - [aws_guardduty_detector_findings](aws_guardduty_detector_findings)
  - [aws_guardduty_detector_intel_sets](aws_guardduty_detector_intel_sets)
  - [aws_guardduty_detector_ip_sets](aws_guardduty_detector_ip_sets)
  - [aws_guardduty_detector_members](aws_guardduty_detector_members)
  - [aws_guardduty_detector_publishing_destinations](aws_guardduty_detector_publishing_destinations)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id (PK)|utf8|
|region (PK)|utf8|
|arn|utf8|
|id (PK)|utf8|
|service_role|utf8|
|status|utf8|
|created_at|timestamp[us, tz=UTC]|
|data_sources|json|
|finding_publishing_frequency|utf8|
|tags|json|
|updated_at|timestamp[us, tz=UTC]|
|result_metadata|json|