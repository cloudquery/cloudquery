
# Table: aws_route53_hosted_zone_query_logging_configs
A complex type that contains information about a configuration for DNS query logging.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|hosted_zone_cq_id|uuid|Unique CloudQuery ID of aws_route53_hosted_zones table (FK)|
|cloud_watch_logs_log_group_arn|text|The Amazon Resource Name (ARN) of the CloudWatch Logs log group that Amazon Route 53 is publishing logs to.|
|id|text|The ID for a configuration for DNS query logging.|
|arn|text|Amazon Resource Name (ARN) of the route53 hosted zone query logging config.|
