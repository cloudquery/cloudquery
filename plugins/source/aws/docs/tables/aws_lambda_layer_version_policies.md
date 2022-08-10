
# Table: aws_lambda_layer_version_policies

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|layer_version_cq_id|uuid|Unique CloudQuery ID of aws_lambda_layer_versions table (FK)|
|layer_version|bigint|The version number.|
|policy|text|The policy document.|
|revision_id|text|A unique identifier for the current revision of the policy.|
