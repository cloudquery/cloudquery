# Table: aws_iam_instance_profiles

This table shows data for IAM Instance Profiles.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_InstanceProfile.html

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|id (PK)|`utf8`|
|tags|`json`|
|arn|`utf8`|
|create_date|`timestamp[us, tz=UTC]`|
|instance_profile_id|`utf8`|
|instance_profile_name|`utf8`|
|path|`utf8`|
|roles|`json`|