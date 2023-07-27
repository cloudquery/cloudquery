# Table: aws_route53recoveryreadiness_recovery_groups

This table shows data for Amazon Route 53 Application Recovery Controller Recovery Readiness Recovery Groups.

https://docs.aws.amazon.com/recovery-readiness/latest/api/recoverygroups.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|cells|`list<item: utf8, nullable>`|
|recovery_group_arn|`utf8`|
|recovery_group_name|`utf8`|
|tags|`json`|