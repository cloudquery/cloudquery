
# Table: aws_apigatewayv2_domain_name_rest_api_mappings
Represents an API mapping.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|domain_name_id|uuid|Unique ID of aws_apigatewayv2_domain_names table (FK)|
|api_id|text|The API identifier.|
|stage|text|The API stage.|
|api_mapping_id|text|The API mapping identifier.|
|api_mapping_key|text|The API mapping key.|
