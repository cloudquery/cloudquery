
# Table: aws_lambda_layer_versions
Details about a version of an AWS Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html). 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|layer_cq_id|uuid|Unique CloudQuery ID of aws_lambda_layers table (FK)|
|compatible_runtimes|text[]|The layer's compatible runtimes.|
|created_date|text|The date that the version was created, in ISO 8601 format|
|description|text|The description of the version.|
|layer_version_arn|text|The ARN of the layer version.|
|license_info|text|The layer's open-source license.|
|version|bigint|The version number.|
