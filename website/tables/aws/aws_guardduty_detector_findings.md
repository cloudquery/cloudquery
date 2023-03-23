# Table: aws_guardduty_detector_findings

This table shows data for Amazon GuardDuty Detector Findings.

https://docs.aws.amazon.com/guardduty/latest/APIReference/API_Finding.html

The primary key for this table is **arn**.

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
|account_id|String|
|arn (PK)|String|
|created_at|String|
|id|String|
|region|String|
|resource|JSON|
|schema_version|String|
|severity|Float|
|type|String|
|updated_at|String|
|confidence|Float|
|description|String|
|partition|String|
|service|JSON|
|title|String|