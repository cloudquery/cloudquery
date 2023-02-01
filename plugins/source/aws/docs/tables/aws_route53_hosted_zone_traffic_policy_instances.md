# Table: aws_route53_hosted_zone_traffic_policy_instances

https://docs.aws.amazon.com/Route53/latest/APIReference/API_TrafficPolicyInstance.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_route53_hosted_zones](aws_route53_hosted_zones.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|arn (PK)|String|
|hosted_zone_arn|String|
|hosted_zone_id|String|
|id|String|
|message|String|
|name|String|
|state|String|
|ttl|Int|
|traffic_policy_id|String|
|traffic_policy_type|String|
|traffic_policy_version|Int|