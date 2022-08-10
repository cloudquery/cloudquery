
# Table: aws_glue_databases
The Database object represents a logical grouping of tables that might reside in a Hive metastore or an RDBMS
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the workflow.|
|name|text|The name of the database|
|catalog_id|text|The ID of the Data Catalog in which the database resides|
|create_table_default_permissions|jsonb|Creates a set of default permissions on the table for principals|
|create_time|timestamp without time zone|The time at which the metadata database was created in the catalog|
|description|text|A description of the database|
|location_uri|text|The location of the database (for example, an HDFS path)|
|parameters|jsonb|These key-value pairs define parameters and properties of the database|
|target_database_catalog_id|text|The ID of the Data Catalog in which the database resides|
|target_database_name|text|The name of the catalog database|
