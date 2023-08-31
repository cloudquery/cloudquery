# Table: aws_quicksight_folders

This table shows data for QuickSight Folders.

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Folder.html

The composite primary key for this table is (**account_id**, **region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|folder_id|`utf8`|
|folder_path|`list<item: utf8, nullable>`|
|folder_type|`utf8`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|name|`utf8`|