
# Table: aws_cloudfront_distribution_custom_error_responses
A complex type that controls:  * Whether CloudFront replaces HTTP status codes in the 4xx and 5xx range with custom error messages before returning the response to the viewer.  * How long CloudFront caches HTTP status codes in the 4xx and 5xx range.  For more information about custom error pages, see Customizing Error Responses (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages.html) in the Amazon CloudFront Developer Guide.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|distribution_cq_id|uuid|Unique CloudQuery ID of aws_cloudfront_distributions table (FK)|
|error_code|integer|The HTTP status code for which you want to specify a custom error page and/or a caching duration.|
|error_caching_min_ttl|bigint|The minimum amount of time, in seconds, that you want CloudFront to cache the HTTP status code specified in ErrorCode|
|response_code|text|The HTTP status code that you want CloudFront to return to the viewer along with the custom error page|
|response_page_path|text|The path to the custom error page that you want CloudFront to return to a viewer when your origin returns the HTTP status code specified by ErrorCode, for example, /4xx-errors/403-forbidden.html|
