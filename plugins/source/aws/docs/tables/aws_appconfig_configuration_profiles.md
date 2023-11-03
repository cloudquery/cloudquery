# Table: aws_appconfig_configuration_profiles

This table shows data for AWS AppConfig Configuration Profiles.

https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_GetConfigurationProfile.html

The composite primary key for this table is (**application_arn**, **arn**).

## Relations

This table depends on [aws_appconfig_applications](aws_appconfig_applications.md).

The following tables depend on aws_appconfig_configuration_profiles:
  - [aws_appconfig_hosted_configuration_versions](aws_appconfig_hosted_configuration_versions.md)

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
|location_uri|`utf8`|
|name|`utf8`|
|retrieval_role_arn|`utf8`|
|type|`utf8`|
|validators|`json`|