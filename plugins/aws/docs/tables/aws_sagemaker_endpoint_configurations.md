
# Table: aws_sagemaker_endpoint_configurations
Provides summary information for an endpoint configuration.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|kms_key_id|text|Amazon Web Services KMS key ID Amazon SageMaker uses to encrypt data when storing it on the ML storage volume attached to the instance.|
|data_capture_config|jsonb||
|tags|jsonb|The tags associated with the model.|
|creation_time|timestamp without time zone|A timestamp that indicates when the endpoint configuration was created.|
|arn|text|The Amazon Resource Name (ARN) of the endpoint configuration.|
|name|text|Name of the Amazon SageMaker endpoint configuration.|
