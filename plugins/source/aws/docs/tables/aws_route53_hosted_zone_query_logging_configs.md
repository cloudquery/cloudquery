# Table: aws_route53_hosted_zone_query_logging_configs

This table shows data for Amazon Route 53 Hosted Zone Query Logging Configs.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_QueryLoggingConfig.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

This table depends on [aws_route53_hosted_zones](aws_route53_hosted_zones.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|hosted_zone_arn|`utf8`|
|cloud_watch_logs_log_group_arn|`utf8`|
|hosted_zone_id|`utf8`|
|id|`utf8`|