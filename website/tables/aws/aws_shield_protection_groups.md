# Table: aws_shield_protection_groups

This table shows data for Shield Protection Groups.

https://docs.aws.amazon.com/waf/latest/DDOSAPIReference/API_ProtectionGroup.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|arn (PK)|String|
|tags|JSON|
|aggregation|String|
|members|StringArray|
|pattern|String|
|protection_group_id|String|
|protection_group_arn|String|
|resource_type|String|