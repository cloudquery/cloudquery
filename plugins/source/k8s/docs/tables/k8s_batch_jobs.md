# Table: k8s_batch_jobs

This table shows data for Kubernetes (K8s) Batch Jobs.

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
|spec_parallelism|`int64`|
|spec_completions|`int64`|
|spec_active_deadline_seconds|`int64`|
|spec_pod_failure_policy|`json`|
|spec_backoff_limit|`int64`|
|spec_selector|`json`|
|spec_manual_selector|`bool`|
|spec_template|`json`|
|spec_ttl_seconds_after_finished|`int64`|
|spec_completion_mode|`utf8`|
|spec_suspend|`bool`|
|status_conditions|`json`|
|status_start_time|`timestamp[us, tz=UTC]`|
|status_completion_time|`timestamp[us, tz=UTC]`|
|status_active|`int64`|
|status_succeeded|`int64`|
|status_failed|`int64`|
|status_completed_indexes|`utf8`|
|status_uncounted_terminated_pods|`json`|
|status_ready|`int64`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Job enforces cpu limits

```sql
WITH
  job_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_batch_jobs
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Job enforces cpu limits' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      job_containers
    WHERE
      job_containers.uid = k8s_batch_jobs.uid
      AND (job_containers.container->'resources'->'limits'->>'cpu') IS NULL
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_batch_jobs;
```

### Job enforces cpu requests

```sql
WITH
  job_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_batch_jobs
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Job enforces cpu requests' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      job_containers
    WHERE
      job_containers.uid = k8s_batch_jobs.uid
      AND (job_containers.container->'resources'->'requests'->>'cpu') IS NULL
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_batch_jobs;
```

### Job enforces memory limit

```sql
WITH
  job_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_batch_jobs
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Job enforces memory limit' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      job_containers
    WHERE
      job_containers.uid = k8s_batch_jobs.uid
      AND (job_containers.container->'resources'->'limits'->>'memory') IS NULL
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_batch_jobs;
```

### Job enforces memory requests

```sql
WITH
  job_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_batch_jobs
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Job enforces memory requests' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      job_containers
    WHERE
      job_containers.uid = k8s_batch_jobs.uid
      AND (job_containers.container->'resources'->'requests'->>'memory') IS NULL
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_batch_jobs;
```

### Job containers privileges disabled

```sql
WITH
  job_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_batch_jobs
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Job containers privileges disabled' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      job_containers
    WHERE
      job_containers.uid = k8s_batch_jobs.uid
      AND job_containers.container->'securityContext'->>'privileged' = 'true'
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_batch_jobs;
```

### Job containers privilege escalation disabled

```sql
WITH
  job_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_batch_jobs
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Job containers privilege escalation disabled' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      job_containers
    WHERE
      job_containers.uid = k8s_batch_jobs.uid
      AND job_containers.container->'securityContext'->>'allowPrivilegeEscalation'
        = 'true'
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_batch_jobs;
```

### Jobs container hostNetwork disabled

```sql
SELECT
  uid AS resource_id,
  'Jobs container hostNetwork disabled' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN spec_template->'spec'->>'hostNetwork' = 'true' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_batch_jobs;
```

### Job containers HostPID and HostIPC sharing disabled

```sql
SELECT
  uid AS resource_id,
  'Job containers HostPID and HostIPC sharing disabled' AS title,
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
  k8s_batch_jobs;
```

### Job containers root file system is read-only

```sql
WITH
  job_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_batch_jobs
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Job containers root file system is read-only' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      job_containers
    WHERE
      job_containers.uid = k8s_batch_jobs.uid
      AND job_containers.container->'securityContext'->>'readOnlyRootFilesystem'
        IS DISTINCT FROM 'true'
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_batch_jobs;
```

### Job containers run as non-root

```sql
WITH
  job_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_batch_jobs
        CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers')
            AS value
    )
SELECT
  uid AS resource_id,
  'Job containers run as non-root' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      job_containers
    WHERE
      job_containers.uid = k8s_batch_jobs.uid
      AND job_containers.container->'securityContext'->>'runAsNonRoot'
        IS DISTINCT FROM 'true'
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_batch_jobs;
```


