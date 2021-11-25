
# Table: aws_sagemaker_models
Provides summary information about a model.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|enable_network_isolation|boolean|If True, no inbound or outbound network calls can be made to or from the model container.|
|execution_role_arn|text|The Amazon Resource Name (ARN) of the IAM role that you specified for the model.|
|inference_execution_config|jsonb|Specifies details of how containers in a multi-container endpoint are called.|
|primary_container|jsonb|The location of the primary inference code, associated artifacts, and custom environment map that the inference code uses when it is deployed in production.|
|tags|jsonb|The tags associated with the model.|
|creation_time|timestamp without time zone|A timestamp that indicates when the model was created.|
|arn|text|The Amazon Resource Name (ARN) of the model.|
|name|text|The name of the model that you want a summary for.|
