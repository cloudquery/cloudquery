# Table: aws_appconfig_applications

This table shows data for AWS AppConfig Applications.

https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_Application.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_appconfig_applications:
  - [aws_appconfig_configuration_profiles](aws_appconfig_configuration_profiles.md)
  - [aws_appconfig_environments](aws_appconfig_environments.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|description|`utf8`|
|id|`utf8`|
|name|`utf8`|