
# Table: aws_apigateway_rest_api_models
Represents the data structure of a method's request or response payload
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_cq_id|uuid|Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)|
|rest_api_id|text|The API's identifier|
|arn|text|The Amazon Resource Name (ARN) for the resource|
|model_template|text||
|content_type|text|The content-type for the model|
|description|text|The description of the model|
|id|text|The identifier for the model resource|
|name|text|The name of the model|
|schema|text|The schema for the model|
