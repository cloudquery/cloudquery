# Table: aws_quicksight_group_members



The composite primary key for this table is (**user_arn**, **group_arn**).

## Relations
This table depends on [aws_quicksight_groups](aws_quicksight_groups.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|user_arn (PK)|String|
|group_arn (PK)|String|
|member_name|String|