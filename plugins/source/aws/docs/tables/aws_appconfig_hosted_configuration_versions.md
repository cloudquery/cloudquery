# Table: aws_appconfig_hosted_configuration_versions

This table shows data for AWS AppConfig Hosted Configuration Versions.

https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_GetHostedConfigurationVersion.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**application_arn**, **arn**).
## Relations

This table depends on [aws_appconfig_configuration_profiles](aws_appconfig_configuration_profiles.md).

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
|configuration_profile_id|`utf8`|
|content|`binary`|
|content_type|`utf8`|
|description|`utf8`|
|kms_key_arn|`utf8`|
|version_label|`utf8`|
|version_number|`int64`|