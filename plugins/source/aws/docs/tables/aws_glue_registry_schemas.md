
# Table: aws_glue_registry_schemas
An object that contains minimal details for a schema
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|registry_cq_id|uuid|Unique CloudQuery ID of aws_glue_registries table (FK)|
|tags|jsonb|Resource tags.|
|compatibility|text|The compatibility mode of the schema.|
|created_time|text|The date and time the schema was created.|
|data_format|text|The data format of the schema definition|
|description|text|A description of schema if specified when created|
|latest_schema_version|bigint|The latest version of the schema associated with the returned schema definition.|
|next_schema_version|bigint|The next version of the schema associated with the returned schema definition.|
|registry_arn|text|The Amazon Resource Name (ARN) of the registry.|
|registry_name|text|The name of the registry.|
|arn|text|The Amazon Resource Name (ARN) of the schema.|
|schema_checkpoint|bigint|The version number of the checkpoint (the last time the compatibility mode was changed).|
|schema_name|text|The name of the schema.|
|schema_status|text|The status of the schema.|
|updated_time|text|The date and time the schema was updated.|
