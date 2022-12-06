# Table: aws_appstream_directory_configs

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_DirectoryConfig.html

The composite primary key for this table is (**account_id**, **region**, **directory_name**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|directory_name (PK)|String|
|certificate_based_auth_properties|JSON|
|created_time|Timestamp|
|organizational_unit_distinguished_names|StringArray|
|service_account_credentials|JSON|