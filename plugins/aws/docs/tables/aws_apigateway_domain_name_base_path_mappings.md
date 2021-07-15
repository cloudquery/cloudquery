
# Table: aws_apigateway_domain_name_base_path_mappings
Represents the base path that callers of the API must provide as part of the URL after the domain name.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|domain_name_cq_id|uuid|Unique CloudQuery ID of aws_apigateway_domain_names table (FK)|
|domain_name|text|The custom domain name as an API host name, for example, my-api.example.com.|
|base_path|text|The base path name that callers of the API must provide as part of the URL after the domain name.|
|rest_api_id|text|The string identifier of the associated RestApi.|
|stage|text|The name of the associated stage.|
