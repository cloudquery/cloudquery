
# Table: aws_apigateway_rest_api_resources
Represents an API resource.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_cq_id|uuid|Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)|
|rest_api_id|text|The API's identifier. This identifier is unique across all of your APIs in API Gateway.|
|id|text|The resource's identifier.|
|parent_id|text|The parent resource's identifier.|
|path|text|The full path for this resource.|
|path_part|text|The last path segment for this resource.|
|resource_methods|jsonb|Gets an API resource's method of a given HTTP verb. The resource methods are a map of methods indexed by methods' HTTP verbs enabled on the resource. This method map is included in the 200 OK response of the GET /restapis/{restapi_id}/resources/{resource_id} or GET /restapis/{restapi_id}/resources/{resource_id}?embed=methods request. Example: Get the GET method of an API resource|
