# Table: aws_route53_hosted_zones



The primary key for this table is **arn**.

## Relations

The following tables depend on aws_route53_hosted_zones:
  - [aws_route53_hosted_zone_query_logging_configs](aws_route53_hosted_zone_query_logging_configs.md)
  - [aws_route53_hosted_zone_resource_record_sets](aws_route53_hosted_zone_resource_record_sets.md)
  - [aws_route53_hosted_zone_traffic_policy_instances](aws_route53_hosted_zone_traffic_policy_instances.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|arn (PK)|String|
|caller_reference|String|
|id|String|
|name|String|
|config|JSON|
|linked_service|JSON|
|resource_record_set_count|Int|
|tags|JSON|
|delegation_set_id|String|
|vpcs|JSON|