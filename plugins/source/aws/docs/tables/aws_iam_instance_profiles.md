# Table: aws_iam_instance_profiles

This table shows data for IAM Instance Profiles.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_InstanceProfile.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|id|`utf8`|
|tags|`json`|
|arn|`utf8`|
|create_date|`timestamp[us, tz=UTC]`|
|instance_profile_id|`utf8`|
|instance_profile_name|`utf8`|
|path|`utf8`|
|roles|`json`|