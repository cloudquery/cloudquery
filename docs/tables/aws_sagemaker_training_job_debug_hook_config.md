
# Table: aws_sagemaker_training_job_debug_hook_config
Configuration information for the Debugger hook parameters, metric and tensor collections, and storage paths
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|training_job_cq_id|uuid|Unique CloudQuery ID of aws_sagemaker_training_jobs table (FK)|
|s3_output_path|text|Path to Amazon S3 storage location for metrics and tensors.  This member is required.|
|collection_configurations|jsonb|Configuration information for Debugger tensor collections|
|hook_parameters|jsonb|Configuration information for the Debugger hook parameters.|
|local_path|text|Path to local storage location for metrics and tensors|
