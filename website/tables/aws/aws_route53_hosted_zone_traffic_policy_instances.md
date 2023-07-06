# Table: aws_route53_hosted_zone_traffic_policy_instances

This table shows data for Amazon Route 53 Hosted Zone Traffic Policy Instances.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_TrafficPolicyInstance.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_route53_hosted_zones](aws_route53_hosted_zones).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|arn (PK)|`utf8`|
|hosted_zone_arn|`utf8`|
|hosted_zone_id|`utf8`|
|id|`utf8`|
|message|`utf8`|
|name|`utf8`|
|state|`utf8`|
|ttl|`int64`|
|traffic_policy_id|`utf8`|
|traffic_policy_type|`utf8`|
|traffic_policy_version|`int64`|