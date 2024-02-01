# Table: aws_route53_traffic_policy_versions

This table shows data for Amazon Route 53 Traffic Policy Versions.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_TrafficPolicy.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**traffic_policy_arn**, **id**, **version**).
## Relations

This table depends on [aws_route53_traffic_policies](aws_route53_traffic_policies.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|traffic_policy_arn|`utf8`|
|id|`utf8`|
|version|`int64`|
|document|`json`|
|name|`utf8`|
|type|`utf8`|
|comment|`utf8`|