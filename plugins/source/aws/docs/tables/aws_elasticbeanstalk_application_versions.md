# Table: aws_elasticbeanstalk_application_versions

This table shows data for AWS Elastic Beanstalk Application Versions.

https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ApplicationVersionDescription.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|application_name|`utf8`|
|application_version_arn|`utf8`|
|build_arn|`utf8`|
|date_created|`timestamp[us, tz=UTC]`|
|date_updated|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|source_build_information|`json`|
|source_bundle|`json`|
|status|`utf8`|
|version_label|`utf8`|