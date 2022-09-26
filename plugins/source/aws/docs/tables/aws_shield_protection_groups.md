# Table: aws_shield_protection_groups


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|arn (PK)|String|
|tags|JSON|
|aggregation|String|
|members|StringArray|
|pattern|String|
|protection_group_id|String|
|resource_type|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|