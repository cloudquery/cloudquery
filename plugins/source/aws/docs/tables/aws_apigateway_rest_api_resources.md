
# Table: aws_apigateway_rest_api_resources
Represents an API resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_cq_id|uuid|Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)|
|rest_api_id|text|The API's identifier|
|arn|text|The Amazon Resource Name (ARN) for the resource|
|id|text|The resource's identifier|
|parent_id|text|The parent resource's identifier|
|path|text|The full path for this resource|
|path_part|text|The last path segment for this resource|
|resource_methods|jsonb|Gets an API resource's method of a given HTTP verb|
