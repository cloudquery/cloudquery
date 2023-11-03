# Table: aws_route53recoverycontrolconfig_routing_controls

This table shows data for Amazon Route 53 Application Recovery Controller Recovery Control Configuration Routing Controls.

https://docs.aws.amazon.com/routing-control/latest/APIReference/API_ListRoutingControls.html

The composite primary key for this table is (**arn**, **control_panel_arn**).

## Relations

This table depends on [aws_route53recoverycontrolconfig_control_panels](aws_route53recoverycontrolconfig_control_panels.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|control_panel_arn (PK)|`utf8`|
|name|`utf8`|
|routing_control_arn|`utf8`|
|status|`utf8`|