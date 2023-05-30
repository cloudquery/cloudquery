# Table: aws_elbv2_load_balancers

This table shows data for Amazon Elastic Load Balancer (ELB) v2 Load Balancers.

https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_LoadBalancer.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_elbv2_load_balancers:
  - [aws_elbv2_listeners](aws_elbv2_listeners)
  - [aws_elbv2_load_balancer_attributes](aws_elbv2_load_balancer_attributes)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|web_acl_arn|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|availability_zones|`json`|
|canonical_hosted_zone_id|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|customer_owned_ipv4_pool|`utf8`|
|dns_name|`utf8`|
|ip_address_type|`utf8`|
|load_balancer_arn|`utf8`|
|load_balancer_name|`utf8`|
|scheme|`utf8`|
|security_groups|`list<item: utf8, nullable>`|
|state|`json`|
|type|`utf8`|
|vpc_id|`utf8`|