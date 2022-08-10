
# Table: aws_glue_database_tables
Represents a collection of related data organized in columns and rows
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|database_cq_id|uuid|Unique CloudQuery ID of aws_glue_databases table (FK)|
|name|text|The table name|
|catalog_id|text|The ID of the Data Catalog in which the table resides|
|create_time|timestamp without time zone|The time when the table definition was created in the Data Catalog|
|created_by|text|The person or entity who created the table|
|database_name|text|The name of the database where the table metadata resides|
|description|text|A description of the table|
|is_registered_with_lake_formation|boolean|Indicates whether the table has been registered with Lake Formation|
|last_access_time|timestamp without time zone|The last time that the table was accessed|
|last_analyzed_time|timestamp without time zone|The last time that column statistics were computed for this table|
|owner|text|The owner of the table|
|parameters|jsonb|These key-value pairs define properties associated with the table|
|retention|bigint|The retention time for this table|
|storage_descriptor|jsonb|A storage descriptor containing information about the physical storage of this table|
|table_type|text|The type of this table (EXTERNAL_TABLE, VIRTUAL_VIEW, etc)|
|target_table_catalog_id|text|The ID of the Data Catalog in which the table resides|
|target_table_database_name|text|The name of the catalog database that contains the target table|
|target_table_name|text|The name of the target table|
|update_time|timestamp without time zone|The last time that the table was updated|
|version_id|text|The ID of the table version|
|view_expanded_text|text|If the table is a view, the expanded text of the view; otherwise null|
|view_original_text|text|If the table is a view, the original text of the view; otherwise null|
