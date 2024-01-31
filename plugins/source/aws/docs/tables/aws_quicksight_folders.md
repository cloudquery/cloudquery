# Table: aws_quicksight_folders

This table shows data for QuickSight Folders.

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Folder.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|folder_id|`utf8`|
|folder_path|`list<item: utf8, nullable>`|
|folder_type|`utf8`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|sharing_model|`utf8`|