# Table: k8s_apps_daemon_sets

This table shows data for Kubernetes (K8s) Apps Daemon Sets.

The primary key for this table is **uid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|context|`utf8`|
|kind|`utf8`|
|api_version|`utf8`|
|name|`utf8`|
|namespace|`utf8`|
|uid (PK)|`utf8`|
|resource_version|`utf8`|
|generation|`int64`|
|deletion_grace_period_seconds|`int64`|
|labels|`json`|
|annotations|`json`|
|owner_references|`json`|
|finalizers|`list<item: utf8, nullable>`|
|spec_selector|`json`|
|spec_template|`json`|
|spec_update_strategy|`json`|
|spec_min_ready_seconds|`int64`|
|spec_revision_history_limit|`int64`|
|status_current_number_scheduled|`int64`|
|status_number_misscheduled|`int64`|
|status_desired_number_scheduled|`int64`|
|status_number_ready|`int64`|
|status_observed_generation|`int64`|
|status_updated_number_scheduled|`int64`|
|status_number_available|`int64`|
|status_number_unavailable|`int64`|
|status_collision_count|`int64`|
|status_conditions|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Daemonset enforces cpu limits

```sql
-- Join every row in the daemonset table with its json array of containers.
WITH
  daemonset_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_daemon_sets
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Daemonset enforces cpu limits' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      daemonset_containers
    WHERE
      daemonset_containers.uid = k8s_apps_daemon_sets.uid
      AND (
          daemonset_containers.container->'resources'->'limits'->>'cpu'
        ) IS NULL
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_daemon_sets;
```

### Daemonset enforces cpu requests

```sql
-- Join every row in the daemonset table with its json array of containers.
WITH
  daemonset_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_daemon_sets
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Daemonset enforces cpu requests' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      daemonset_containers
    WHERE
      daemonset_containers.uid = k8s_apps_daemon_sets.uid
      AND (
          daemonset_containers.container->'resources'->'requests'->>'cpu'
        ) IS NULL
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_daemon_sets;
```

### Daemonset enforces memory limits

```sql
-- Join every row in the daemonset table with its json array of containers.
WITH
  daemonset_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_daemon_sets
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Daemonset enforces memory limits' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      daemonset_containers
    WHERE
      daemonset_containers.uid = k8s_apps_daemon_sets.uid
      AND (
          daemonset_containers.container->'resources'->'limits'->>'memory'
        ) IS NULL
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_daemon_sets;
```

### Daemonset enforces memory requests

```sql
-- Join every row in the daemonset table with its json array of containers.
WITH
  daemonset_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_daemon_sets
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Daemonset enforces memory requests' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      daemonset_containers
    WHERE
      daemonset_containers.uid = k8s_apps_daemon_sets.uid
      AND (
          daemonset_containers.container->'resources'->'requests'->>'memory'
        ) IS NULL
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_daemon_sets;
```

### DaemonSet containers privileges disabled

```sql
WITH
  daemonset_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_daemon_sets
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'DaemonSet containers privileges disabled' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      daemonset_containers
    WHERE
      daemonset_containers.uid = k8s_apps_daemon_sets.uid
      AND daemonset_containers.container->'securityContext'->>'privileged'
        = 'true'
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_daemon_sets;
```

### DaemonSet containers privilege escalation disabled

```sql
WITH
  daemonset_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_daemon_sets
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'DaemonSet containers privilege escalation disabled' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      daemonset_containers
    WHERE
      daemonset_containers.uid = k8s_apps_daemon_sets.uid
      AND daemonset_containers.container->'securityContext'->>'allowPrivilegeEscalation'
        = 'true'
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_daemon_sets;
```

### Deamonset container hostNetwork disabled

```sql
SELECT
  uid AS resource_id,
  'Deamonset container hostNetwork disabled' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN spec_template->'spec'->>'hostNetwork' = 'true' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_daemon_sets;
```

### Deamonset containers HostPID and HostIPC sharing disabled

```sql
SELECT
  uid AS resource_id,
  'Deamonset containers HostPID and HostIPC sharing disabled' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN spec_template->'spec'->>'hostPID' = 'true'
  OR spec_template->'spec'->>'hostIPC' = 'true'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_daemon_sets;
```

### DeamonSet containers root file system is read-only

```sql
WITH
  daemonset_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_daemon_sets
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'DeamonSet containers root file system is read-only' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      daemonset_containers
    WHERE
      daemonset_containers.uid = k8s_apps_daemon_sets.uid
      AND daemonset_containers.container->'securityContext'->>'readOnlyRootFilesystem'
        IS DISTINCT FROM 'true'
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_daemon_sets;
```

### DaemonSet containers to run as non-root

```sql
WITH
  daemonset_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_daemon_sets
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'DaemonSet containers to run as non-root' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      daemonset_containers
    WHERE
      daemonset_containers.uid = k8s_apps_daemon_sets.uid
      AND daemonset_containers.container->'securityContext'->>'runAsNonRoot'
        IS DISTINCT FROM 'true'
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_daemon_sets;
```


