
# Table: aws_apigateway_rest_api_deployments
An immutable representation of a RestApi resource that can be called by users using Stages.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_cq_id|uuid|Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)|
|rest_api_id|text|The API's identifier. This identifier is unique across all of your APIs in API Gateway.|
|arn|text|The Amazon Resource Name (ARN) for the resource.|
|api_summary|jsonb|A summary of the RestApi at the date and time that the deployment resource was created.|
|created_date|timestamp without time zone|The date and time that the deployment resource was created.|
|description|text|The description for the deployment resource.|
|id|text|The identifier for the deployment resource.|
