
# Table: aws_accounts
Information about IAM entity usage and IAM quotas in the AWS account.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account|
|users|integer|Current number of users.|
|users_quota|integer|Maximum allowed number of users.|
|groups|integer|Current number of groups.|
|groups_quota|integer|Maximum allowed number of groups.|
|server_certificates|integer|Current number of server certificates|
|server_certificates_quota|integer|Maximum allowed number of server certificates.|
|user_policy_size_quota|integer|Maximum allowed policies per user|
|group_policy_size_quota|integer|Maximum allowed policies per group|
|groups_per_user_quota|integer|Maximum allowed groups per user|
|signing_certificates_per_user_quota|integer|Maximum allowed server certificates per user|
|access_keys_per_user_quota|integer|Maximum allowed quota of access keys per user.|
|mfa_devices|integer|Number of MFa devices|
|mfa_devices_in_use|integer|Number of MFA devices in use.|
|account_mfa_enabled|boolean|Whether MFA is enabled for the account.|
|account_access_keys_present|boolean|Number of account level access keys present.|
|account_signing_certificates_present|boolean|Number of account signing certificates present.|
|attached_policies_per_group_quota|integer|Maximum allowed attached policies per group.|
|policies|integer|Current number of policies.|
|policies_quota|integer|Allowed number of policies.|
|policy_size_quota|integer|Allowed size of policies.|
|policy_versions_in_use|integer|Number of policy versions in use.|
|policy_versions_in_use_quota|integer|Allowed number of policy versions.|
|versions_per_policy_quota|integer| Allowed number of versions per policy.|
|global_endpoint_token_version|integer|Token version of the global endpoint.|
|aliases|text[]|List of aliases associated with the account. AWS supports only one alias per account|
