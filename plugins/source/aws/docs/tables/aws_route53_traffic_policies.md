# Table: aws_route53_traffic_policies


The primary key for this table is **arn**.

## Relations
The following tables depend on `aws_route53_traffic_policies`:
  - [`aws_route53_traffic_policy_versions`](aws_route53_traffic_policy_versions.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|arn (PK)|String|
|id|String|
|latest_version|Int|
|name|String|
|traffic_policy_count|Int|
|type|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|