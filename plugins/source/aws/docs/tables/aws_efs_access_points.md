# Table: aws_efs_access_points

This table shows data for Amazon Elastic File System (EFS) Access Points.

https://docs.aws.amazon.com/efs/latest/ug/API_AccessPointDescription.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|access_point_arn|`utf8`|
|access_point_id|`utf8`|
|client_token|`utf8`|
|file_system_id|`utf8`|
|life_cycle_state|`utf8`|
|name|`utf8`|
|owner_id|`utf8`|
|posix_user|`json`|
|root_directory|`json`|