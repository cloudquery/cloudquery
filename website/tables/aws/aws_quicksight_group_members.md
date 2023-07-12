# Table: aws_quicksight_group_members

This table shows data for QuickSight Group Members.

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_GroupMember.html

The composite primary key for this table is (**account_id**, **region**, **group_arn**, **arn**).

## Relations

This table depends on [aws_quicksight_groups](aws_quicksight_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|group_arn (PK)|`utf8`|
|arn (PK)|`utf8`|
|member_name|`utf8`|