# Table: gcp_compute_projects

This table shows data for GCP Compute Projects.

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|common_instance_metadata|`json`|
|creation_timestamp|`utf8`|
|default_network_tier|`utf8`|
|default_service_account|`utf8`|
|description|`utf8`|
|enabled_features|`list<item: utf8, nullable>`|
|id|`int64`|
|kind|`utf8`|
|name|`utf8`|
|quotas|`json`|
|self_link (PK)|`utf8`|
|usage_export_location|`json`|
|vm_dns_setting|`utf8`|
|xpn_project_status|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that instances are not configured to use the default service account (Automated)

```sql
SELECT
  DISTINCT
  gci.name AS resource_id,
  'Ensure that instances are not configured to use the default service account (Automated)'
    AS title,
  gci.project_id AS project_id,
  CASE
  WHEN gci.name NOT LIKE 'gke-'
  AND gcisa->>'email'
    = (
        SELECT
          default_service_account
        FROM
          gcp_compute_projects
        WHERE
          project_id = gci.project_id
      )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_compute_instances AS gci,
  jsonb_array_elements(gci.service_accounts) AS gcisa;
```

### Ensure that instances are not configured to use the default service account with full access to all Cloud APIs (Automated)

```sql
SELECT
  DISTINCT
  gci.name AS resource_id,
  'Ensure that instances are not configured to use the default service account with full access to all Cloud APIs (Automated)'
    AS title,
  gci.project_id AS project_id,
  CASE
  WHEN gcisa->>'email'
  = (
      SELECT
        default_service_account
      FROM
        gcp_compute_projects
      WHERE
        project_id = gci.project_id
    )
  AND ARRAY['https://www.googleapis.com/auth/cloud-platform']
    <@ ARRAY (SELECT jsonb_array_elements_text(gcisa->'scopes'))
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_compute_instances AS gci,
  jsonb_array_elements(gci.service_accounts) AS gcisa;
```

### Ensure oslogin is enabled for a Project (Automated)

```sql
SELECT
  name AS resource_id,
  'Ensure oslogin is enabled for a Project (Automated)' AS title,
  project_id AS project_id,
  CASE
  WHEN (cimd->>'key') IS NULL
  OR NOT (cimd->>'value' = ANY ('{1,true,True,TRUE,y,yes}'))
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_compute_projects
  LEFT JOIN jsonb_array_elements(common_instance_metadata->'items') AS cimd ON
      cimd->>'key' = 'enable-oslogin';
```


