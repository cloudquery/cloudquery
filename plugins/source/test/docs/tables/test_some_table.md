# Table: test_some_table

Test table

The composite primary key for this table is (**resource_id**, **client_id**).

## Relations

The following tables depend on test_some_table:
  - [test_sub_table](test_sub_table.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|resource_id (PK)|`int64`|
|column2|`utf8`|
|client_id (PK)|`int64`|