
# Table: aws_apigateway_rest_api_request_validators
A set of validation rules for incoming Method requests.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_cq_id|uuid|Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)|
|rest_api_id|text|The API's identifier. This identifier is unique across all of your APIs in API Gateway.|
|arn|text|The Amazon Resource Name (ARN) for the resource.|
|id|text|The identifier of this RequestValidator.|
|name|text|The name of this RequestValidator|
|validate_request_body|boolean|A Boolean flag to indicate whether to validate a request body according to the configured Model schema.|
|validate_request_parameters|boolean|A Boolean flag to indicate whether to validate request parameters (true) or not (false).|
