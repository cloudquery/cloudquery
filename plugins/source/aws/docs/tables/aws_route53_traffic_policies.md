# Table: aws_route53_traffic_policies

https://docs.aws.amazon.com/Route53/latest/APIReference/API_TrafficPolicySummary.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_route53_traffic_policies:
  - [aws_route53_traffic_policy_versions](aws_route53_traffic_policy_versions.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|arn (PK)|String|
|id|String|
|latest_version|Int|
|name|String|
|traffic_policy_count|Int|
|type|String|