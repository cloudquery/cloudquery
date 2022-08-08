
# Table: aws_glue_ml_transform_input_record_tables
The database and table in the Glue Data Catalog that is used for input or output data
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|ml_transform_cq_id|uuid|Unique CloudQuery ID of aws_glue_ml_transforms table (FK)|
|database_name|text|A database name in the Glue Data Catalog|
|table_name|text|A table name in the Glue Data Catalog|
|catalog_id|text|A unique identifier for the Glue Data Catalog|
|connection_name|text|The name of the connection to the Glue Data Catalog|
