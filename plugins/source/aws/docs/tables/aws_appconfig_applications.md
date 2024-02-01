# Table: aws_appconfig_applications

This table shows data for AWS AppConfig Applications.

https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_Application.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_appconfig_applications:
  - [aws_appconfig_configuration_profiles](aws_appconfig_configuration_profiles.md)
  - [aws_appconfig_environments](aws_appconfig_environments.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|description|`utf8`|
|id|`utf8`|
|name|`utf8`|