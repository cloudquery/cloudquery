# Table: aws_appconfig_configuration_profiles

This table shows data for AWS AppConfig Configuration Profiles.

https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_GetConfigurationProfile.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**application_arn**, **arn**).
## Relations

This table depends on [aws_appconfig_applications](aws_appconfig_applications.md).

The following tables depend on aws_appconfig_configuration_profiles:
  - [aws_appconfig_hosted_configuration_versions](aws_appconfig_hosted_configuration_versions.md)

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
|kms_key_arn|`utf8`|
|kms_key_identifier|`utf8`|
|location_uri|`utf8`|
|name|`utf8`|
|retrieval_role_arn|`utf8`|
|type|`utf8`|
|validators|`json`|