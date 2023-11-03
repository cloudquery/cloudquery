# Table: aws_appconfig_hosted_configuration_versions

This table shows data for AWS AppConfig Hosted Configuration Versions.

https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_GetHostedConfigurationVersion.html

The composite primary key for this table is (**application_arn**, **arn**).

## Relations

This table depends on [aws_appconfig_configuration_profiles](aws_appconfig_configuration_profiles.md).

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
|configuration_profile_id|`utf8`|
|content|`binary`|
|content_type|`utf8`|
|description|`utf8`|
|version_label|`utf8`|
|version_number|`int64`|