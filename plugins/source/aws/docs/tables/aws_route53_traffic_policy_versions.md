# Table: aws_route53_traffic_policy_versions


The composite primary key for this table is (**traffic_policy_arn**, **id**, **version**).

## Relations
This table depends on [`aws_route53_traffic_policies`](aws_route53_traffic_policies.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|traffic_policy_arn (PK)|String|
|id (PK)|String|
|version (PK)|Int|
|document|String|
|name|String|
|type|String|
|comment|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|