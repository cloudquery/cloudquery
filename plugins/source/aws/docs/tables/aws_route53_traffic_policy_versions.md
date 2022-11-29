# Table: aws_route53_traffic_policy_versions

https://docs.aws.amazon.com/Route53/latest/APIReference/API_TrafficPolicy.html

The composite primary key for this table is (**traffic_policy_arn**, **id**, **version**).

## Relations
This table depends on [aws_route53_traffic_policies](aws_route53_traffic_policies.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|traffic_policy_arn (PK)|String|
|id (PK)|String|
|version (PK)|Int|
|document|JSON|
|name|String|
|type|String|
|comment|String|