# Table: aws_glue_security_configurations

This table shows data for Glue Security Configurations.

https://docs.aws.amazon.com/glue/latest/webapi/API_SecurityConfiguration.html

The composite primary key for this table is (**account_id**, **region**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|name (PK)|`utf8`|
|created_time_stamp|`timestamp[us, tz=UTC]`|
|encryption_configuration|`json`|