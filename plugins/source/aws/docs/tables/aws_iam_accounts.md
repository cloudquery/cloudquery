# Table: aws_iam_accounts

This table shows data for IAM Accounts.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetAccountSummary.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **account_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|users|`int64`|
|users_quota|`int64`|
|groups|`int64`|
|groups_quota|`int64`|
|server_certificates|`int64`|
|server_certificates_quota|`int64`|
|user_policy_size_quota|`int64`|
|group_policy_size_quota|`int64`|
|groups_per_user_quota|`int64`|
|signing_certificates_per_user_quota|`int64`|
|access_keys_per_user_quota|`int64`|
|mfa_devices|`int64`|
|mfa_devices_in_use|`int64`|
|account_mfa_enabled|`bool`|
|account_access_keys_present|`bool`|
|account_signing_certificates_present|`bool`|
|attached_policies_per_group_quota|`int64`|
|attached_policies_per_role_quota|`int64`|
|attached_policies_per_user_quota|`int64`|
|policies|`int64`|
|policies_quota|`int64`|
|policy_size_quota|`int64`|
|policy_versions_in_use|`int64`|
|policy_versions_in_use_quota|`int64`|
|versions_per_policy_quota|`int64`|
|global_endpoint_token_version|`int64`|
|aliases|`list<item: utf8, nullable>`|