# Table: aws_iam_accounts



The primary key for this table is **account_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|users|Int|
|users_quota|Int|
|groups|Int|
|groups_quota|Int|
|server_certificates|Int|
|server_certificates_quota|Int|
|user_policy_size_quota|Int|
|group_policy_size_quota|Int|
|groups_per_user_quota|Int|
|signing_certificates_per_user_quota|Int|
|access_keys_per_user_quota|Int|
|mfa_devices|Int|
|mfa_devices_in_use|Int|
|account_mfa_enabled|Bool|
|account_access_keys_present|Bool|
|account_signing_certificates_present|Bool|
|attached_policies_per_group_quota|Int|
|attached_policies_per_role_quota|Int|
|attached_policies_per_user_quota|Int|
|policies|Int|
|policies_quota|Int|
|policy_size_quota|Int|
|policy_versions_in_use|Int|
|policy_versions_in_use_quota|Int|
|versions_per_policy_quota|Int|
|global_endpoint_token_version|Int|
|aliases|StringArray|