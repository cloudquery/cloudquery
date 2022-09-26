# Table: aws_sagemaker_endpoint_configurations


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|creation_time|Timestamp|
|endpoint_config_name|String|
|production_variants|JSON|
|async_inference_config|JSON|
|data_capture_config|JSON|
|kms_key_id|String|
|result_metadata|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|