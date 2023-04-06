# Table: aws_elbv2_listeners

This table shows data for Amazon Elastic Load Balancer (ELB) v2 Listeners.

https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_Listener.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_elbv2_load_balancers](aws_elbv2_load_balancers).

The following tables depend on aws_elbv2_listeners:
  - [aws_elbv2_listener_certificates](aws_elbv2_listener_certificates)
  - [aws_elbv2_listener_rules](aws_elbv2_listener_rules)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|alpn_policy|StringArray|
|certificates|JSON|
|default_actions|JSON|
|listener_arn|String|
|load_balancer_arn|String|
|port|Int|
|protocol|String|
|ssl_policy|String|