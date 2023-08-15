# Table: test_table

This table shows data for Test Table.

Description for test table

The composite primary key for this table is (**id_col**, **id_col2**).

## Relations

The following tables depend on test_table:
  - [relation_table](relation_table.md)
  - [relation_table2](relation_table2.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|int_col|`int64`|
|id_col (PK)|`int64`|
|id_col2 (PK)|`int64`|
|json_col|`json`|
|list_col|`list<item: int64, nullable>`|
|map_col|`map<utf8, int64, items_nullable>`|
|struct_col|`struct<string_field: utf8, int_field: int64>`|
