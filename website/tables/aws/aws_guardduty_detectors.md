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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|arn|String|
|id (PK)|String|
|service_role|String|
|status|String|
|created_at|Timestamp|
|data_sources|JSON|
|finding_publishing_frequency|String|
|tags|JSON|
|updated_at|Timestamp|
|result_metadata|JSON|