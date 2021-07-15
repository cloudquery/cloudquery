
# Table: aws_apigateway_rest_api_documentation_versions
A snapshot of the documentation of an API.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_cq_id|uuid|Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)|
|rest_api_id|text|The API's identifier. This identifier is unique across all of your APIs in API Gateway.|
|created_date|timestamp without time zone|The date when the API documentation snapshot is created.|
|description|text|The description of the API documentation snapshot.|
|version|text|The version identifier of the API documentation snapshot.|
