# Table: aws_signer_signing_profiles

This table shows data for AWS Signer Signing Profiles.

https://docs.aws.amazon.com/signer/latest/api/API_GetSigningProfile.html

The primary key for this table is **profile_version_arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|overrides|`json`|
|platform_display_name|`utf8`|
|platform_id|`utf8`|
|profile_name|`utf8`|
|profile_version|`utf8`|
|profile_version_arn (PK)|`utf8`|
|revocation_record|`json`|
|signature_validity_period|`json`|
|signing_material|`json`|
|signing_parameters|`json`|
|status|`utf8`|
|status_reason|`utf8`|
|tags|`json`|