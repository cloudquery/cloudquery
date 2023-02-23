# Table: aws_quicksight_groups

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Group.html

The composite primary key for this table is (**account_id**, **region**, **arn**).

## Relations

The following tables depend on aws_quicksight_groups:
  - [aws_quicksight_group_members](aws_quicksight_group_members.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|tags|JSON|
|arn (PK)|String|
|description|String|
|group_name|String|
|principal_id|String|