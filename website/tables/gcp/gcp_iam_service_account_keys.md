# Table: gcp_iam_service_account_keys

This table shows data for GCP IAM Service Account Keys.

https://cloud.google.com/iam/docs/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKey

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_iam_service_accounts](gcp_iam_service_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|service_account_unique_id|`utf8`|
|name (PK)|`utf8`|
|key_algorithm|`utf8`|
|public_key_data|`binary`|
|valid_after_time|`timestamp[us, tz=UTC]`|
|valid_before_time|`timestamp[us, tz=UTC]`|
|key_origin|`utf8`|
|key_type|`utf8`|
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


