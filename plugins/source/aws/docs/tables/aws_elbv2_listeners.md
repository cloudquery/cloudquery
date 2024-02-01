# Table: aws_elbv2_listeners

This table shows data for Amazon Elastic Load Balancer (ELB) v2 Listeners.

https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_Listener.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

This table depends on [aws_elbv2_load_balancers](aws_elbv2_load_balancers.md).

The following tables depend on aws_elbv2_listeners:
  - [aws_elbv2_listener_certificates](aws_elbv2_listener_certificates.md)
  - [aws_elbv2_listener_rules](aws_elbv2_listener_rules.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|alpn_policy|`list<item: utf8, nullable>`|
|certificates|`json`|
|default_actions|`json`|
|listener_arn|`utf8`|
|load_balancer_arn|`utf8`|
|mutual_authentication|`json`|
|port|`int64`|
|protocol|`utf8`|
|ssl_policy|`utf8`|