# Table: gcp_iam_service_accounts

This table shows data for GCP IAM Service Accounts.

https://cloud.google.com/iam/docs/reference/rest/v1/projects.serviceAccounts#ServiceAccount

The composite primary key for this table is (**unique_id**, **name**).

## Relations

The following tables depend on gcp_iam_service_accounts:
  - [gcp_iam_service_account_keys](gcp_iam_service_account_keys)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|unique_id (PK)|`utf8`|
|name (PK)|`utf8`|
|email|`utf8`|
|display_name|`utf8`|
|etag|`binary`|
|description|`utf8`|
|oauth2_client_id|`utf8`|
|disabled|`bool`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that there are only GCP-managed service account keys for each service account (Automated)

```sql
SELECT
  DISTINCT
  gisa.name AS resource_id,
  'Ensure that there are only GCP-managed service account keys for each service account (Automated)'
    AS title,
  gisa.project_id AS project_id,
  CASE
  WHEN gisa.email LIKE '%iam.gserviceaccount.com'
  AND gisak.key_type = 'USER_MANAGED'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_iam_service_accounts AS gisa
  JOIN gcp_iam_service_account_keys AS gisak ON
      gisa.project_id = gisak.project_id
      AND gisa.unique_id = gisak.service_account_unique_id;
```

### Ensure user-managed/external keys for service accounts are rotated every 90 days or less (Automated)

```sql
SELECT
  DISTINCT
  gisa.name AS resource_id,
  'Ensure user-managed/external keys for service accounts are rotated every 90 days or less (Automated)'
    AS title,
  gisa.project_id AS project_id,
  CASE
  WHEN gisa.email LIKE '%iam.gserviceaccount.com'
  AND gisak.valid_after_time::TIMESTAMP <= (now() - '90'::INTERVAL DAY)
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_iam_service_accounts AS gisa
  JOIN gcp_iam_service_account_keys AS gisak ON
      gisa.project_id = gisak.project_id
      AND gisa.unique_id = gisak.service_account_unique_id;
```


