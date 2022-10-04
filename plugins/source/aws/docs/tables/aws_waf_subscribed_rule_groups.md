# Table: aws_waf_subscribed_rule_groups



The composite primary key for this table is (**account_id**, **rule_group_id**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id (PK)|String|
|rule_group_id (PK)|String|
|metric_name|String|
|name|String|