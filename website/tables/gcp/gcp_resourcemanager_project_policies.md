# Table: gcp_resourcemanager_project_policies

This table shows data for GCP Resourcemanager Project Policies.

https://cloud.google.com/resource-manager/reference/rest/Shared.Types/Policy

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|audit_configs|`json`|
|bindings|`json`|
|etag|`utf8`|
|version|`int64`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that Cloud Audit Logging is configured properly across all services and all users from a project (Automated)

```sql
WITH
  project_policy_audit_configs
    AS (
      SELECT
        project_id, jsonb_array_elements(audit_configs) AS audit_config
      FROM
        gcp_resourcemanager_project_policies
      WHERE
        audit_configs != 'null'
    ),
  log_types
    AS (
      SELECT
        project_id,
        audit_config->>'service' AS service,
        jsonb_array_elements(audit_config->'auditLogConfigs')->>'logType'
          AS logs,
        jsonb_array_elements(
          audit_config->'auditLogConfigs'
        )->>'exemptedMembers'
          AS exempted
      FROM
        project_policy_audit_configs
    ),
  valid_log_types
    AS (
      SELECT
        project_id, service, count(*) AS valid_types
      FROM
        log_types
      WHERE
        exempted IS NULL
        AND logs IN ('ADMIN_READ', 'DATA_READ', 'DATA_WRITE')
        AND service = 'allServices'
      GROUP BY
        project_id, service
    )
SELECT
  service AS resource_id,
  'Ensure that Cloud Audit Logging is configured properly across all services and all users from a project (Automated)'
    AS title,
  project_id AS project_id,
  CASE WHEN valid_types = 3 THEN 'pass' ELSE 'fail' END AS status
FROM
  valid_log_types;
```


