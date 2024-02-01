# Table: aws_route53recoverycontrolconfig_routing_controls

This table shows data for Amazon Route 53 Application Recovery Controller Recovery Control Configuration Routing Controls.

https://docs.aws.amazon.com/routing-control/latest/APIReference/API_ListRoutingControls.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**arn**, **control_panel_arn**).
## Relations

This table depends on [aws_route53recoverycontrolconfig_control_panels](aws_route53recoverycontrolconfig_control_panels.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|control_panel_arn|`utf8`|
|name|`utf8`|
|owner|`utf8`|
|routing_control_arn|`utf8`|
|status|`utf8`|