# Table: aws_elbv1_load_balancers

This table shows data for Amazon Elastic Load Balancer (ELB) v1 Load Balancers.

https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_LoadBalancerDescription.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_elbv1_load_balancers:
  - [aws_elbv1_load_balancer_policies](aws_elbv1_load_balancer_policies)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id|utf8|
|region|utf8|
|arn (PK)|utf8|
|availability_zones|list<item: utf8, nullable>|
|backend_server_descriptions|json|
|canonical_hosted_zone_name|utf8|
|canonical_hosted_zone_name_id|utf8|
|created_time|timestamp[us, tz=UTC]|
|dns_name|utf8|
|health_check|json|
|instances|json|
|listener_descriptions|json|
|load_balancer_name|utf8|
|policies|json|
|scheme|utf8|
|security_groups|list<item: utf8, nullable>|
|source_security_group|json|
|subnets|list<item: utf8, nullable>|
|vpc_id|utf8|
|tags|json|
|attributes|json|