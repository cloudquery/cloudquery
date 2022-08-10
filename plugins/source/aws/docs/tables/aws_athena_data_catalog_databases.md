
# Table: aws_athena_data_catalog_databases
Contains metadata information for a database in a data catalog
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|data_catalog_cq_id|uuid|Unique CloudQuery ID of aws_athena_data_catalogs table (FK)|
|name|text|The name of the database|
|description|text|An optional description of the database|
|parameters|jsonb|A set of custom key/value pairs|
