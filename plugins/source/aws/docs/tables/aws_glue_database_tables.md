
# Table: aws_glue_database_tables
Represents a collection of related data organized in columns and rows
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|database_cq_id|uuid|Unique CloudQuery ID of aws_glue_databases table (FK)|
|parameters|jsonb||
|storage_parameters|jsonb||
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
|retention|bigint|The retention time for this table|
|additional_locations|text[]|A list of locations that point to the path where a Delta table is located|
|bucket_columns|text[]|A list of reducer grouping columns, clustering columns, and bucketing columns in the table|
|compressed|boolean|True if the data in the table is compressed, or False if not|
|input_format|text|The input format: SequenceFileInputFormat (binary), or TextInputFormat, or a custom format|
|location|text|The physical location of the table|
|number_of_buckets|bigint|Must be specified if the table contains any dimension columns|
|output_format|text|The output format: SequenceFileOutputFormat (binary), or IgnoreKeyTextOutputFormat, or a custom format|
|schema_reference_schema_id|jsonb|A structure that contains schema identity fields|
|schema_reference_schema_version_id|text|The unique ID assigned to a version of the schema|
|schema_reference_schema_version_number|bigint|The version number of the schema|
|serde_info|jsonb|The serialization/deserialization (SerDe) information|
|skewed_info|jsonb|The information about values that appear frequently in a column (skewed values)|
|sort_columns|jsonb|A list specifying the sort order of each bucket in the table|
|stored_as_sub_directories|boolean|True if the table data is stored in subdirectories, or False if not|
|table_type|text|The type of this table (EXTERNAL_TABLE, VIRTUAL_VIEW, etc)|
|target_table_catalog_id|text|The ID of the Data Catalog in which the table resides|
|target_table_database_name|text|The name of the catalog database that contains the target table|
|target_table_name|text|The name of the target table|
|update_time|timestamp without time zone|The last time that the table was updated|
|version_id|text|The ID of the table version|
|view_expanded_text|text|If the table is a view, the expanded text of the view; otherwise null|
|view_original_text|text|If the table is a view, the original text of the view; otherwise null|
