
# Table: aws_athena_data_catalogs
Contains information about a data catalog in an Amazon Web Services account
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|ARN of the resource.|
|tags|jsonb|Tags associated with the Athena data catalog.|
|name|text|The name of the data catalog|
|type|text|The type of data catalog to create: LAMBDA for a federated catalog, HIVE for an external hive metastore, or GLUE for an Glue Data Catalog|
|description|text|An optional description of the data catalog|
|parameters|jsonb|Specifies the Lambda function or functions to use for the data catalog|
