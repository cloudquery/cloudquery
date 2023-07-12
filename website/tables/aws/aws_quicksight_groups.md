# Table: aws_quicksight_groups

This table shows data for QuickSight Groups.

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Group.html

The composite primary key for this table is (**account_id**, **region**, **arn**).

## Relations

The following tables depend on aws_quicksight_groups:
  - [aws_quicksight_group_members](aws_quicksight_group_members)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|description|`utf8`|
|group_name|`utf8`|
|principal_id|`utf8`|