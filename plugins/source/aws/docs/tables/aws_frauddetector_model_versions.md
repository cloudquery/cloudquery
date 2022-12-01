# Table: aws_frauddetector_model_versions

https://docs.aws.amazon.com/frauddetector/latest/api/API_ModelVersionDetail.html

The primary key for this table is **arn**.

## Relations
This table depends on [aws_frauddetector_models](aws_frauddetector_models.md).


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
|external_events_detail|JSON|
|ingested_events_detail|JSON|
|last_updated_time|String|
|model_id|String|
|model_type|String|
|model_version_number|String|
|status|String|
|training_data_schema|JSON|
|training_data_source|String|
|training_result|JSON|
|training_result_v2|JSON|