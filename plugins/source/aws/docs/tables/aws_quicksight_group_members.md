# Table: aws_quicksight_group_members

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_GroupMember.html

The composite primary key for this table is (**account_id**, **region**, **group_arn**, **arn**).

## Relations

This table depends on [aws_quicksight_groups](aws_quicksight_groups.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|group_arn (PK)|String|
|arn (PK)|String|
|member_name|String|