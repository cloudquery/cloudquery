# Table: aws_elbv2_load_balancers

This table shows data for Amazon Elastic Load Balancer (ELB) v2 Load Balancers.

https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_LoadBalancer.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_elbv2_load_balancers:
  - [aws_elbv2_listeners](aws_elbv2_listeners.md)
  - [aws_elbv2_load_balancer_attributes](aws_elbv2_load_balancer_attributes.md)
  - [aws_elbv2_load_balancer_web_acls](aws_elbv2_load_balancer_web_acls.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|availability_zones|`json`|
|canonical_hosted_zone_id|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|customer_owned_ipv4_pool|`utf8`|
|dns_name|`utf8`|
|enforce_security_group_inbound_rules_on_private_link_traffic|`utf8`|
|ip_address_type|`utf8`|
|load_balancer_arn|`utf8`|
|load_balancer_name|`utf8`|
|scheme|`utf8`|
|security_groups|`list<item: utf8, nullable>`|
|state|`json`|
|type|`utf8`|
|vpc_id|`utf8`|