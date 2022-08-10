
# Table: aws_glue_registry_schema_versions
An object containing the details about a schema version
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|registry_schema_cq_id|uuid|Unique CloudQuery ID of aws_glue_registry_schemas table (FK)|
|metadata|jsonb||
|created_time|text|The date and time the schema version was created.|
|data_format|text|The data format of the schema definition|
|schema_arn|text|The Amazon Resource Name (ARN) of the schema.|
|schema_definition|text|The schema definition for the schema ID.|
|id|text|The SchemaVersionId of the schema version.|
|status|text|The status of the schema version.|
|version_number|bigint|The version number of the schema.|
