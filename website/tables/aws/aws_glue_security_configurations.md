# Table: aws_glue_security_configurations

This table shows data for Glue Security Configurations.

https://docs.aws.amazon.com/glue/latest/webapi/API_SecurityConfiguration.html

The composite primary key for this table is (**account_id**, **region**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|name (PK)|String|
|created_time_stamp|Timestamp|
|encryption_configuration|JSON|