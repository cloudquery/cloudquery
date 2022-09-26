# Table: aws_sagemaker_models


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|creation_time|Timestamp|
|execution_role_arn|String|
|model_name|String|
|containers|JSON|
|enable_network_isolation|Bool|
|inference_execution_config|JSON|
|primary_container|JSON|
|vpc_config|JSON|
|result_metadata|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|