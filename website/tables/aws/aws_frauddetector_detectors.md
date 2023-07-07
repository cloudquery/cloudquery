# Table: aws_frauddetector_detectors

This table shows data for Amazon Fraud Detector Detectors.

https://docs.aws.amazon.com/frauddetector/latest/api/API_Detector.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_frauddetector_detectors:
  - [aws_frauddetector_rules](aws_frauddetector_rules)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|created_time|`utf8`|
|description|`utf8`|
|detector_id|`utf8`|
|event_type_name|`utf8`|
|last_updated_time|`utf8`|