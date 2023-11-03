# Table: aws_sagemaker_endpoint_configurations

This table shows data for Amazon SageMaker Endpoint Configurations.

https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeEndpointConfig.html

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
|endpoint_config_arn|`utf8`|
|endpoint_config_name|`utf8`|
|production_variants|`json`|
|async_inference_config|`json`|
|data_capture_config|`json`|
|explainer_config|`json`|
|kms_key_id|`utf8`|
|shadow_production_variants|`json`|
|result_metadata|`json`|