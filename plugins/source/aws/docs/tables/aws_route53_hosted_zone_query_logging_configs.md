# Table: aws_route53_hosted_zone_query_logging_configs


The primary key for this table is **arn**.

## Relations
This table depends on [`aws_route53_hosted_zones`](aws_route53_hosted_zones.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|arn (PK)|String|
|hosted_zone_arn|String|
|cloud_watch_logs_log_group_arn|String|
|hosted_zone_id|String|
|id|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|