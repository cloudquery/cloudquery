# Table: aws_wafv2_managed_rule_groups


The composite primary key for this table is (**account_id**, **region**, **scope**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id (PK)|String|
|region (PK)|String|
|scope (PK)|String|
|properties|JSON|
|description|String|
|name|String|
|vendor_name|String|
|versioning_supported|Bool|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|