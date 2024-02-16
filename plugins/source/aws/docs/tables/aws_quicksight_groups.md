# Table: aws_quicksight_groups

This table shows data for QuickSight Groups.

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Group.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **arn**).
## Relations

The following tables depend on aws_quicksight_groups:
  - [aws_quicksight_group_members](aws_quicksight_group_members.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|description|`utf8`|
|group_name|`utf8`|
|principal_id|`utf8`|