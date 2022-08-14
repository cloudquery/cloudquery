
# Table: aws_lightsail_database_snapshots
Describes a database snapshot
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the database snapshot|
|created_at|timestamp without time zone|The timestamp when the database snapshot was created|
|engine|text|The software of the database snapshot (for example, MySQL)|
|engine_version|text|The database engine version for the database snapshot (for example, 5723)|
|from_relational_database_arn|text|The Amazon Resource Name (ARN) of the database from which the database snapshot was created|
|from_relational_database_blueprint_id|text|The blueprint ID of the database from which the database snapshot was created|
|from_relational_database_bundle_id|text|The bundle ID of the database from which the database snapshot was created|
|from_relational_database_name|text|The name of the source database from which the database snapshot was created|
|availability_zone|text|The Availability Zone|
|name|text|The name of the database snapshot|
|resource_type|text|The Lightsail resource type|
|size_in_gb|bigint|The size of the disk in GB (for example, 32) for the database snapshot|
|state|text|The state of the database snapshot|
|support_code|text|The support code for the database snapshot|
|tags|jsonb|The tag keys and optional values for the resource|
