# Table: aws_appstream_directory_configs

This table shows data for Amazon AppStream Directory Configs.

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_DirectoryConfig.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **directory_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|directory_name|`utf8`|
|certificate_based_auth_properties|`json`|
|created_time|`timestamp[us, tz=UTC]`|
|organizational_unit_distinguished_names|`list<item: utf8, nullable>`|
|service_account_credentials|`json`|