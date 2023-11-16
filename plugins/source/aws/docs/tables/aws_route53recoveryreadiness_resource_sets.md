# Table: aws_route53recoveryreadiness_resource_sets

This table shows data for Amazon Route 53 Application Recovery Controller Recovery Readiness Resource Sets.

https://docs.aws.amazon.com/recovery-readiness/latest/api/resourcesets.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|resource_set_arn|`utf8`|
|resource_set_name|`utf8`|
|resource_set_type|`utf8`|
|resources|`json`|
|tags|`json`|