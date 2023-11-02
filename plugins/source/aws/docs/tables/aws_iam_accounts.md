# Table: aws_iam_accounts

This table shows data for IAM Accounts.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetAccountSummary.html

The primary key for this table is **account_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
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

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Security contact information should be provided for an AWS account

```sql
SELECT
  'Security contact information should be provided for an AWS account' AS title,
  aws_iam_accounts.account_id,
  CASE WHEN alternate_contact_type IS NULL THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_iam_accounts
  LEFT JOIN (
      SELECT
        *
      FROM
        aws_account_alternate_contacts
      WHERE
        alternate_contact_type = 'SECURITY'
    )
      AS account_security_contacts ON
      aws_iam_accounts.account_id = account_security_contacts.account_id;
```

### S3 Block Public Access setting should be enabled

```sql
SELECT
  'S3 Block Public Access setting should be enabled' AS title,
  aws_iam_accounts.account_id,
  aws_iam_accounts.account_id AS resource_id,
  CASE
  WHEN config_exists IS NOT true
  OR block_public_acls IS NOT true
  OR block_public_policy IS NOT true
  OR ignore_public_acls IS NOT true
  OR restrict_public_buckets IS NOT true
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_iam_accounts
  LEFT JOIN aws_s3_accounts ON
      aws_iam_accounts.account_id = aws_s3_accounts.account_id;
```

### SSM documents should not be public

```sql
SELECT
  'SSM documents should not be public' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN 'all' = ANY (ARRAY (SELECT jsonb_array_elements_text(p->'AccountIds')))
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_ssm_documents, jsonb_array_elements(aws_ssm_documents.permissions) AS p
WHERE
  owner IN (SELECT account_id FROM aws_iam_accounts);
```


