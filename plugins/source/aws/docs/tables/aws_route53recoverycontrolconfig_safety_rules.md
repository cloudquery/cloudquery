# Table: aws_route53recoverycontrolconfig_safety_rules

This table shows data for Amazon Route 53 Application Recovery Controller Recovery Control Configuration Safety Rules.

https://docs.aws.amazon.com/recovery-cluster/latest/api/controlpanel-controlpanelarn-safetyrules.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

This table depends on [aws_route53recoverycontrolconfig_control_panels](aws_route53recoverycontrolconfig_control_panels.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|assertion|`json`|
|gating|`json`|