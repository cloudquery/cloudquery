# Table: aws_appconfig_environments

This table shows data for AWS AppConfig Environments.

https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_Environment.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**application_arn**, **arn**).
## Relations

This table depends on [aws_appconfig_applications](aws_appconfig_applications.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|application_arn|`utf8`|
|arn|`utf8`|
|application_id|`utf8`|
|description|`utf8`|
|id|`utf8`|
|monitors|`json`|
|name|`utf8`|
|state|`utf8`|