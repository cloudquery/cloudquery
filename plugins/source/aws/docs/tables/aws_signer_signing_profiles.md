# Table: aws_signer_signing_profiles

This table shows data for AWS Signer Signing Profiles.

https://docs.aws.amazon.com/signer/latest/api/API_GetSigningProfile.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **profile_version_arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|overrides|`json`|
|platform_display_name|`utf8`|
|platform_id|`utf8`|
|profile_name|`utf8`|
|profile_version|`utf8`|
|profile_version_arn|`utf8`|
|revocation_record|`json`|
|signature_validity_period|`json`|
|signing_material|`json`|
|signing_parameters|`json`|
|status|`utf8`|
|status_reason|`utf8`|
|tags|`json`|