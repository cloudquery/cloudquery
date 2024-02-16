# Table: test_sub_table

Sub table of test_some_table

The composite primary key for this table is (**parent_resource_id**, **sub_resource_id**).

## Relations

This table depends on [test_some_table](test_some_table.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|parent_resource_id (PK)|`int64`|
|sub_resource_id (PK)|`int64`|
|data_column|`utf8`|