# Table: k8s_apps_replica_sets

This table shows data for Kubernetes (K8s) Apps Replica Sets.

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
|spec_min_ready_seconds|`int64`|
|spec_selector|`json`|
|spec_template|`json`|
|status_replicas|`int64`|
|status_fully_labeled_replicas|`int64`|
|status_ready_replicas|`int64`|
|status_available_replicas|`int64`|
|status_observed_generation|`int64`|
|status_conditions|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Replicaset enforces cpu limits

```sql
-- Join every row in the replica_set table with its json array of containers.
WITH
  replica_set_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_replica_sets
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Replicaset enforces cpu limits' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      replica_set_containers
    WHERE
      replica_set_containers.uid = k8s_apps_replica_sets.uid
      AND (
          replica_set_containers.container->'resources'->'limits'->>'cpu'
        ) IS NULL
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_replica_sets;
```

### Replicaset enforces cpu requests

```sql
-- Join every row in the replica_set table with its json array of containers.
WITH
  replica_set_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_replica_sets
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Replicaset enforces cpu requests' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      replica_set_containers
    WHERE
      replica_set_containers.uid = k8s_apps_replica_sets.uid
      AND (
          replica_set_containers.container->'resources'->'requests'->>'cpu'
        ) IS NULL
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_replica_sets;
```

### Replicaset enforces memory limits

```sql
-- Join every row in the replica_set table with its json array of containers.
WITH
  replica_set_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_replica_sets
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Replicaset enforces memory limits' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      replica_set_containers
    WHERE
      replica_set_containers.uid = k8s_apps_replica_sets.uid
      AND (
          replica_set_containers.container->'resources'->'limits'->>'memory'
        ) IS NULL
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_replica_sets;
```

### Replicaset enforces memory requests

```sql
-- Join every row in the deployment table with its json array of containers.
WITH
  replica_set_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_replica_sets
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Replicaset enforces memory requests' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      replica_set_containers
    WHERE
      replica_set_containers.uid = k8s_apps_replica_sets.uid
      AND (
          replica_set_containers.container->'resources'->'requests'->>'memory'
        ) IS NULL
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_replica_sets;
```

### Replicaset privileges disabled

```sql
WITH
  replica_set_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_replica_sets
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Replicaset privileges disabled' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      replica_set_containers
    WHERE
      replica_set_containers.uid = k8s_apps_replica_sets.uid
      AND replica_set_containers.container->'securityContext'->>'privileged'
        = 'true'
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_replica_sets;
```

### ReplicaSet container privileged escalation disabled

```sql
WITH
  replica_set_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_replica_sets
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'ReplicaSet container privileged escalation disabled' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      replica_set_containers
    WHERE
      replica_set_containers.uid = k8s_apps_replica_sets.uid
      AND replica_set_containers.container->'securityContext'->>'privileged'
        = 'true'
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_replica_sets;
```

### ReplicaSet container hostNetwork disabled

```sql
SELECT
  uid AS resource_id,
  'ReplicaSet container hostNetwork disabled' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN spec_template->'spec'->>'hostNetwork' = 'true' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_replica_sets;
```

### ReplicaSet containers HostPID and HostIPC sharing disabled

```sql
SELECT
  uid AS resource_id,
  'ReplicaSet containers HostPID and HostIPC sharing disabled' AS title,
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
  k8s_apps_replica_sets;
```

### ReplicaSet containers root file system is read-only

```sql
WITH
  replica_set_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_replica_sets
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'ReplicaSet containers root file system is read-only' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      replica_set_containers
    WHERE
      replica_set_containers.uid = k8s_apps_replica_sets.uid
      AND replica_set_containers.container->'securityContext'->>'readOnlyRootFilesystem'
        IS DISTINCT FROM 'true'
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_replica_sets;
```

### ReplicaSet containers must run as non-root

```sql
WITH
  replica_set_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_apps_replica_sets
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'ReplicaSet containers must run as non-root' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      replica_set_containers
    WHERE
      replica_set_containers.uid = k8s_apps_replica_sets.uid
      AND replica_set_containers.container->'securityContext'->>'runAsNonRoot'
        IS DISTINCT FROM 'true'
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_apps_replica_sets;
```


