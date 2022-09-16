
# Table: aws_lambda_layers
Details about an AWS Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html). 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|latest_matching_version_compatible_runtimes|text[]|The layer's compatible runtimes.|
|latest_matching_version_created_date|timestamp without time zone|The date that the version was created, in ISO 8601 format|
|latest_matching_version_description|text|The description of the version.|
|latest_matching_version_layer_version_arn|text|The ARN of the layer version.|
|latest_matching_version_license_info|text|The layer's open-source license.|
|latest_matching_version|bigint|The version number.|
|arn|text|The Amazon Resource Name (ARN) of the function layer.|
|name|text|The name of the layer.|
