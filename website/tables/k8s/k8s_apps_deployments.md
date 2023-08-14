# Table: k8s_apps_deployments

This table shows data for Kubernetes (K8s) Apps Deployments.

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
|spec_replicas|`int64`|
|spec_selector|`json`|
|spec_template|`json`|
|spec_strategy|`json`|
|spec_min_ready_seconds|`int64`|
|spec_revision_history_limit|`int64`|
|spec_paused|`bool`|
|spec_progress_deadline_seconds|`int64`|
|status_observed_generation|`int64`|
|status_replicas|`int64`|
|status_updated_replicas|`int64`|
|status_ready_replicas|`int64`|
|status_available_replicas|`int64`|
|status_unavailable_replicas|`int64`|
|status_conditions|`json`|
|status_collision_count|`int64`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Deployment enforces cpu limits

```sql
-- Join every row in the deployment table with its json array of containers.
WITH
  deployment_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_deployments
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Deployment enforces cpu limits' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      deployment_containers
    WHERE
      deployment_containers.uid = k8s_apps_deployments.uid
      AND (
          deployment_containers.container->'resources'->'limits'->>'cpu'
        ) IS NULL
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_deployments;
```

### Deployment enforces cpu requests

```sql
-- Join every row in the deployment table with its json array of containers.
WITH
  deployment_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_deployments
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Deployment enforces cpu requests' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      deployment_containers
    WHERE
      deployment_containers.uid = k8s_apps_deployments.uid
      AND (
          deployment_containers.container->'resources'->'requests'->>'cpu'
        ) IS NULL
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_deployments;
```

### Deployment enforces memory limits

```sql
-- Join every row in the deployment table with its json array of containers.
WITH
  deployment_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_deployments
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Deployment enforces memory limits' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      deployment_containers
    WHERE
      deployment_containers.uid = k8s_apps_deployments.uid
      AND (
          deployment_containers.container->'resources'->'limits'->>'memory'
        ) IS NULL
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_deployments;
```

### Deployment enforces memory requests

```sql
-- Join every row in the deployment table with its json array of containers.
WITH
  deployment_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_deployments
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Deployment enforces memory requests' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      deployment_containers
    WHERE
      deployment_containers.uid = k8s_apps_deployments.uid
      AND (
          deployment_containers.container->'resources'->'requests'->>'memory'
        ) IS NULL
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_deployments;
```

### Deployments privileges disabled

```sql
WITH
  deployment_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_deployments
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Deployments privileges disabled' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      deployment_containers
    WHERE
      deployment_containers.uid = k8s_apps_deployments.uid
      AND deployment_containers.container->'securityContext'->>'privileged'
        = 'true'
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_deployments;
```

### Deployments privilege escalation disabled

```sql
WITH
  deployment_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_deployments
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Deployments privilege escalation disabled' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      deployment_containers
    WHERE
      deployment_containers.uid = k8s_apps_deployments.uid
      AND deployment_containers.container->'securityContext'->>'allowPrivilegeEscalation'
        = 'true'
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_deployments;
```

### Deployments container hostNetwork disabled

```sql
SELECT
  uid AS resource_id,
  'Deployments container hostNetwork disabled' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN spec_template->'spec'->>'hostNetwork' = 'true' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_deployments;
```

### Deployment containers HostPID and HostIPC sharing disabled

```sql
SELECT
  uid AS resource_id,
  'Deployment containers HostPID and HostIPC sharing disabled' AS title,
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
  k8s_apps_deployments;
```

### Deployment containers root file system is read-only

```sql
WITH
  deployment_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_deployments
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Deployment containers root file system is read-only' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      deployment_containers
    WHERE
      deployment_containers.uid = k8s_apps_deployments.uid
      AND deployment_containers.container->'securityContext'->>'readOnlyRootFilesystem'
        IS DISTINCT FROM 'true'
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_deployments;
```

### Deployment containers to run as non-root

```sql
WITH
  deployment_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_deployments
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Deployment containers to run as non-root' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      deployment_containers
    WHERE
      deployment_containers.uid = k8s_apps_deployments.uid
      AND deployment_containers.container->'securityContext'->>'runAsNonRoot'
        IS DISTINCT FROM 'true'
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_deployments;
```


