
# Table: aws_lambda_function_layers
An Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_cq_id|uuid|Unique CloudQuery ID of aws_lambda_functions table (FK)|
|function_arn|text|The Amazon Resource Name (ARN) of the lambda function|
|arn|text|The Amazon Resource Name (ARN) of the function layer.|
|code_size|bigint|The size of the layer archive in bytes.|
|signing_job_arn|text|The Amazon Resource Name (ARN) of a signing job.|
|signing_profile_version_arn|text|The Amazon Resource Name (ARN) for a signing profile version.|
