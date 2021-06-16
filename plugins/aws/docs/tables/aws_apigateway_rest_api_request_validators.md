
# Table: aws_apigateway_rest_api_request_validators
A set of validation rules for incoming Method requests.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_id|uuid|Unique ID of aws_apigateway_rest_apis table (FK)|
|resource_id|text|The identifier of this RequestValidator.|
|name|text|The name of this RequestValidator|
|validate_request_body|boolean|A Boolean flag to indicate whether to validate a request body according to the configured Model schema.|
|validate_request_parameters|boolean|A Boolean flag to indicate whether to validate request parameters (true) or not (false).|
