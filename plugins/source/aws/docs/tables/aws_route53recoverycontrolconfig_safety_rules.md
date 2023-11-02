# Table: aws_route53recoverycontrolconfig_safety_rules

This table shows data for Amazon Route 53 Application Recovery Controller Recovery Control Configuration Safety Rules.

https://docs.aws.amazon.com/recovery-cluster/latest/api/controlpanel-controlpanelarn-safetyrules.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_route53recoverycontrolconfig_control_panels](aws_route53recoverycontrolconfig_control_panels.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|assertion|`json`|
|gating|`json`|