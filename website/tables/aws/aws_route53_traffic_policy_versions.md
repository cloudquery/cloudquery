# Table: aws_route53_traffic_policy_versions

This table shows data for Amazon Route 53 Traffic Policy Versions.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_TrafficPolicy.html

The composite primary key for this table is (**traffic_policy_arn**, **id**, **version**).

## Relations

This table depends on [aws_route53_traffic_policies](aws_route53_traffic_policies).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|traffic_policy_arn (PK)|`utf8`|
|id (PK)|`utf8`|
|version (PK)|`int64`|
|document|`json`|
|name|`utf8`|
|type|`utf8`|
|comment|`utf8`|