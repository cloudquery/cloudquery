# Table: aws_emr_studios

This table shows data for Amazon EMR Studios.

https://docs.aws.amazon.com/emr/latest/APIReference/API_Studio.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_emr_studios:
  - [aws_emr_studio_session_mappings](aws_emr_studio_session_mappings.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|auth_mode|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|default_s3_location|`utf8`|
|description|`utf8`|
|encryption_key_arn|`utf8`|
|engine_security_group_id|`utf8`|
|idc_instance_arn|`utf8`|
|idc_user_assignment|`utf8`|
|idp_auth_url|`utf8`|
|idp_relay_state_parameter_name|`utf8`|
|name|`utf8`|
|service_role|`utf8`|
|studio_arn|`utf8`|
|studio_id|`utf8`|
|subnet_ids|`list<item: utf8, nullable>`|
|tags|`json`|
|trusted_identity_propagation_enabled|`bool`|
|url|`utf8`|
|user_role|`utf8`|
|vpc_id|`utf8`|
|workspace_security_group_id|`utf8`|