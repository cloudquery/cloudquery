# Table: aws_frauddetector_models

https://docs.aws.amazon.com/frauddetector/latest/api/API_Model.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_frauddetector_models:
  - [aws_frauddetector_model_versions](aws_frauddetector_model_versions.md)

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
|created_time|String|
|description|String|
|event_type_name|String|
|last_updated_time|String|
|model_id|String|
|model_type|String|