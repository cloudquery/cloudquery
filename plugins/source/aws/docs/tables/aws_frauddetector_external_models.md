# Table: aws_frauddetector_external_models

This table shows data for Amazon Fraud Detector External Models.

https://docs.aws.amazon.com/frauddetector/latest/api/API_ExternalModel.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|created_time|`utf8`|
|input_configuration|`json`|
|invoke_model_endpoint_role_arn|`utf8`|
|last_updated_time|`utf8`|
|model_endpoint|`utf8`|
|model_endpoint_status|`utf8`|
|model_source|`utf8`|
|output_configuration|`json`|