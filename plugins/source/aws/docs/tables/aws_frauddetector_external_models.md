# Table: aws_frauddetector_external_models

https://docs.aws.amazon.com/frauddetector/latest/api/API_ExternalModel.html

The primary key for this table is **arn**.



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
|input_configuration|JSON|
|invoke_model_endpoint_role_arn|String|
|last_updated_time|String|
|model_endpoint|String|
|model_endpoint_status|String|
|model_source|String|
|output_configuration|JSON|