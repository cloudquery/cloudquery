# Table: aws_appconfig_environments

This table shows data for AWS AppConfig Environments.

https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_Environment.html

The composite primary key for this table is (**application_arn**, **arn**).

## Relations

This table depends on [aws_appconfig_applications](aws_appconfig_applications.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|application_arn (PK)|`utf8`|
|arn (PK)|`utf8`|
|application_id|`utf8`|
|description|`utf8`|
|id|`utf8`|
|monitors|`json`|
|name|`utf8`|
|state|`utf8`|