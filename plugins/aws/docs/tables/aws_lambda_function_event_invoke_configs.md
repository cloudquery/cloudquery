
# Table: aws_lambda_function_event_invoke_configs
A configuration object that specifies the destination of an event after Lambda processes it. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_cq_id|uuid|Unique CloudQuery ID of aws_lambda_functions table (FK)|
|on_failure_destination|text|The Amazon Resource Name (ARN) of the destination resource.|
|on_success_destination|text|The Amazon Resource Name (ARN) of the destination resource.|
|function_arn|text|The Amazon Resource Name (ARN) of the function.|
|last_modified|timestamp without time zone|The date and time that the configuration was last updated.|
|maximum_event_age_in_seconds|integer|The maximum age of a request that Lambda sends to a function for processing.|
|maximum_retry_attempts|integer|The maximum number of times to retry when the function returns an error.|
