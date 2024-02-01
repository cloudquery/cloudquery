# Table: aws_quicksight_group_members

This table shows data for QuickSight Group Members.

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_GroupMember.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **group_arn**, **arn**).
## Relations

This table depends on [aws_quicksight_groups](aws_quicksight_groups.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|group_arn|`utf8`|
|arn|`utf8`|
|member_name|`utf8`|