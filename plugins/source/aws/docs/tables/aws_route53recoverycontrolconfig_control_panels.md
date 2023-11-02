# Table: aws_route53recoverycontrolconfig_control_panels

This table shows data for Amazon Route 53 Application Recovery Controller Recovery Control Configuration Control Panels.

https://docs.aws.amazon.com/recovery-cluster/latest/api/controlpanels.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_route53recoverycontrolconfig_control_panels:
  - [aws_route53recoverycontrolconfig_routing_controls](aws_route53recoverycontrolconfig_routing_controls.md)
  - [aws_route53recoverycontrolconfig_safety_rules](aws_route53recoverycontrolconfig_safety_rules.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|cluster_arn|`utf8`|
|control_panel_arn|`utf8`|
|default_control_panel|`bool`|
|name|`utf8`|
|routing_control_count|`int64`|
|status|`utf8`|