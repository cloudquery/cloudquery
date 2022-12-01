# Table: aws_elbv2_load_balancers

https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_LoadBalancer.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_elbv2_load_balancers:
  - [aws_elbv2_listeners](aws_elbv2_listeners.md)
  - [aws_elbv2_load_balancer_attributes](aws_elbv2_load_balancer_attributes.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|web_acl_arn|String|
|tags|JSON|
|arn (PK)|String|
|availability_zones|JSON|
|canonical_hosted_zone_id|String|
|created_time|Timestamp|
|customer_owned_ipv4_pool|String|
|dns_name|String|
|ip_address_type|String|
|load_balancer_name|String|
|scheme|String|
|security_groups|StringArray|
|state|JSON|
|type|String|
|vpc_id|String|