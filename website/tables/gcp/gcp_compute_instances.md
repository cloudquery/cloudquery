# Table: gcp_compute_instances

This table shows data for GCP Compute Instances.

https://cloud.google.com/compute/docs/reference/rest/v1/instances#Instance

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|advanced_machine_features|`json`|
|can_ip_forward|`bool`|
|confidential_instance_config|`json`|
|cpu_platform|`utf8`|
|creation_timestamp|`utf8`|
|deletion_protection|`bool`|
|description|`utf8`|
|disks|`json`|
|display_device|`json`|
|fingerprint|`utf8`|
|guest_accelerators|`json`|
|hostname|`utf8`|
|id|`int64`|
|key_revocation_action_type|`utf8`|
|kind|`utf8`|
|label_fingerprint|`utf8`|
|labels|`json`|
|last_start_timestamp|`utf8`|
|last_stop_timestamp|`utf8`|
|last_suspended_timestamp|`utf8`|
|machine_type|`utf8`|
|metadata|`json`|
|min_cpu_platform|`utf8`|
|name|`utf8`|
|network_interfaces|`json`|
|network_performance_config|`json`|
|params|`json`|
|private_ipv6_google_access|`utf8`|
|reservation_affinity|`json`|
|resource_policies|`list<item: utf8, nullable>`|
|resource_status|`json`|
|satisfies_pzs|`bool`|
|scheduling|`json`|
|self_link (PK)|`utf8`|
|service_accounts|`json`|
|shielded_instance_config|`json`|
|shielded_instance_integrity_policy|`json`|
|source_machine_image|`utf8`|
|source_machine_image_encryption_key|`json`|
|start_restricted|`bool`|
|status|`utf8`|
|status_message|`utf8`|
|tags|`json`|
|zone|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that IP forwarding is not enabled on Instances (Automated)

```sql
SELECT
  name AS resource_id,
  'Ensure that IP forwarding is not enabled on Instances (Automated)' AS title,
  project_id AS project_id,
  CASE WHEN can_ip_forward = true THEN 'fail' ELSE 'pass' END AS status
FROM
  gcp_compute_instances;
```

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

### Ensure that Compute instances do not have public IP addresses (Automated

```sql
SELECT
  DISTINCT
  gci.name AS resource_id,
  'Ensure that Compute instances do not have public IP addresses (Automated'
    AS title,
  gci.project_id AS project_id,
  CASE
  WHEN gci.name NOT LIKE 'gke-%'
  AND (
      (ac4->>'nat_i_p') IS NOT NULL
      OR ac4->>'nat_i_p' != ''
      OR (ac6->>'nat_i_p') IS NOT NULL
      OR ac6->>'nat_i_p' != ''
    )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_compute_instances AS gci,
  jsonb_array_elements(gci.network_interfaces) AS ni
  LEFT JOIN jsonb_array_elements(ni->'access_configs') AS ac4 ON true
  LEFT JOIN jsonb_array_elements(ni->'ipv6_access_configs') AS ac6 ON true;
```

### Ensure Compute instances are launched with Shielded VM enabled (Automated)

```sql
SELECT
  name AS resource_id,
  'Ensure Compute instances are launched with Shielded VM enabled (Automated)'
    AS title,
  project_id AS project_id,
  CASE
  WHEN (shielded_instance_config->>'enable_integrity_monitoring')::BOOL = false
  OR (shielded_instance_config->>'enable_vtpm')::BOOL = false
  OR (shielded_instance_config->>'enable_secure_boot')::BOOL = false
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_compute_instances;
```

### Ensure "Block Project-wide SSH keys" is enabled for VM instances (Automated)

```sql
SELECT
  gci.name AS resource_id,
  'Ensure "Block Project-wide SSH keys" is enabled for VM instances (Automated)'
    AS title,
  gci.project_id AS project_id,
  CASE
  WHEN (gcmi->>'key') IS NULL
  OR NOT (gcmi->>'value' = ANY ('{1,true,True,TRUE,y,yes}'))
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_compute_instances AS gci
  LEFT JOIN jsonb_array_elements(gci.metadata->'items') AS gcmi ON
      gcmi->>'key' = 'block-project-ssh-keys';
```

### Ensure that Compute instances have Confidential Computing enabled (Automated)

```sql
SELECT
  name AS resource_id,
  'Ensure that Compute instances have Confidential Computing enabled (Automated)'
    AS title,
  project_id AS project_id,
  CASE
  WHEN (confidential_instance_config->>'enable_confidential_compute')::BOOL
  = false
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_compute_instances;
```

### Ensure "Enable connecting to serial ports" is not enabled for VM Instance (Automated)

```sql
SELECT
  name AS resource_id,
  'Ensure "Enable connecting to serial ports" is not enabled for VM Instance (Automated)'
    AS title,
  project_id AS project_id,
  CASE
  WHEN gcmi->>'key' = 'serial-port-enable'
  AND gcmi->>'value' = ANY ('{1,true,True,TRUE,y,yes}')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_compute_instances AS gci,
  jsonb_array_elements(gci.metadata->'items') AS gcmi;
```


