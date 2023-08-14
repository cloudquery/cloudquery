# Table: aws_frauddetector_model_versions

This table shows data for Amazon Fraud Detector Model Versions.

https://docs.aws.amazon.com/frauddetector/latest/api/API_ModelVersionDetail.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_frauddetector_models](aws_frauddetector_models).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|created_time|`utf8`|
|external_events_detail|`json`|
|ingested_events_detail|`json`|
|last_updated_time|`utf8`|
|model_id|`utf8`|
|model_type|`utf8`|
|model_version_number|`utf8`|
|status|`utf8`|
|training_data_schema|`json`|
|training_data_source|`utf8`|
|training_result|`json`|
|training_result_v2|`json`|