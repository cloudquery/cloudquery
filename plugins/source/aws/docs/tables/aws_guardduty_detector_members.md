# Table: aws_guardduty_detector_members


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_guardduty_detectors`](aws_guardduty_detectors.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|region|String|
|detector_arn|String|
|account_id|String|
|email|String|
|master_id|String|
|relationship_status|String|
|updated_at|String|
|administrator_id|String|
|detector_id|String|
|invited_at|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|