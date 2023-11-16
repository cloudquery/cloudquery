# Table: aws_iot_security_profiles

This table shows data for AWS IoT Security Profiles.

https://docs.aws.amazon.com/iot/latest/apireference/API_DescribeSecurityProfile.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|targets|`list<item: utf8, nullable>`|
|tags|`json`|
|arn (PK)|`utf8`|
|additional_metrics_to_retain|`list<item: utf8, nullable>`|
|additional_metrics_to_retain_v2|`json`|
|alert_targets|`json`|
|behaviors|`json`|
|creation_date|`timestamp[us, tz=UTC]`|
|last_modified_date|`timestamp[us, tz=UTC]`|
|security_profile_arn|`utf8`|
|security_profile_description|`utf8`|
|security_profile_name|`utf8`|
|version|`int64`|
|result_metadata|`json`|