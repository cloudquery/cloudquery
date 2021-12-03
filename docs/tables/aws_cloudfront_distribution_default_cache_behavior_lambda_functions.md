
# Table: aws_cloudfront_distribution_default_cache_behavior_lambda_functions
A complex type that contains a Lambda function association.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|distribution_cq_id|uuid|Unique CloudQuery ID of aws_cloudfront_distributions table (FK)|
|event_type|text|Specifies the event type that triggers a Lambda function invocation|
|lambda_function_arn|text|The ARN of the Lambda function|
|include_body|boolean|A flag that allows a Lambda function to have read access to the body content. For more information, see Accessing the Request Body by Choosing the Include Body Option (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-include-body-access.html) in the Amazon CloudFront Developer Guide.|
