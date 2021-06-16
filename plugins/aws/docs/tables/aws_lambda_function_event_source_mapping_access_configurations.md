
# Table: aws_lambda_function_event_source_mapping_access_configurations
You can specify the authentication protocol, or the VPC components to secure access to your event source.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_event_source_mapping_id|uuid|Unique ID of aws_lambda_function_event_source_mappings table (FK)|
|type|text|The type of authentication protocol or the VPC components for your event source.|
|uri|text|The value for your chosen configuration in Type|
