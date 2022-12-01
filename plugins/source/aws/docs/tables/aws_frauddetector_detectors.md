# Table: aws_frauddetector_detectors

https://docs.aws.amazon.com/frauddetector/latest/api/API_Detector.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_frauddetector_detectors:
  - [aws_frauddetector_rules](aws_frauddetector_rules.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|created_time|String|
|description|String|
|detector_id|String|
|event_type_name|String|
|last_updated_time|String|