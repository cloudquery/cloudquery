# Table: aws_route53_hosted_zone_traffic_policy_instances

This table shows data for Amazon Route 53 Hosted Zone Traffic Policy Instances.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_TrafficPolicyInstance.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **arn**).
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
|hosted_zone_id|`utf8`|
|id|`utf8`|
|message|`utf8`|
|name|`utf8`|
|state|`utf8`|
|ttl|`int64`|
|traffic_policy_id|`utf8`|
|traffic_policy_type|`utf8`|
|traffic_policy_version|`int64`|