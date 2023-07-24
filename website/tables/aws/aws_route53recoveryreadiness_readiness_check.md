# Table: aws_route53recoveryreadiness_readiness_check

This table shows data for Amazon Route 53 Application Recovery Controller Recovery Readiness Readiness Check.

https://docs.aws.amazon.com/recovery-readiness/latest/api/readinesschecks.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|readiness_check_arn|`utf8`|
|resource_set|`utf8`|
|readiness_check_name|`utf8`|
|tags|`json`|