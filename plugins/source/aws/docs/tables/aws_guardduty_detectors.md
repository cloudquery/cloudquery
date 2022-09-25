# Table: aws_guardduty_detectors


The composite primary key for this table is (**account_id**, **region**, **id**).

## Relations
The following tables depend on `aws_guardduty_detectors`:
  - [`aws_guardduty_detector_members`](aws_guardduty_detector_members.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id (PK)|String|
|region (PK)|String|
|arn|String|
|id (PK)|String|
|service_role|String|
|status|String|
|created_at|String|
|data_sources|JSON|
|finding_publishing_frequency|String|
|tags|JSON|
|updated_at|String|
|result_metadata|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|