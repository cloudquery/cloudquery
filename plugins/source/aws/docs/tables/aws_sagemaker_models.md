# Table: aws_sagemaker_models

This table shows data for Amazon SageMaker Models.

https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeModel.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|creation_time|`timestamp[us, tz=UTC]`|
|execution_role_arn|`utf8`|
|model_arn|`utf8`|
|model_name|`utf8`|
|containers|`json`|
|deployment_recommendation|`json`|
|enable_network_isolation|`bool`|
|inference_execution_config|`json`|
|primary_container|`json`|
|vpc_config|`json`|
|result_metadata|`json`|