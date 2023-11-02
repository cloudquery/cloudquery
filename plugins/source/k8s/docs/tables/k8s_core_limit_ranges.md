# Table: k8s_core_limit_ranges

This table shows data for Kubernetes (K8s) Core Limit Ranges.

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
|spec_limits|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Namespaces CPU default resource limit

```sql
WITH
  default_cpu_limits
    AS (
      SELECT
        context, namespace, value->'default'->>'cpu' AS default_cpu_limit
      FROM
        k8s_core_limit_ranges
        CROSS JOIN jsonb_array_elements(k8s_core_limit_ranges.spec_limits)
    )
INSERT
INTO
  k8s_policy_results
    (
      resource_id,
      execution_time,
      framework,
      check_id,
      title,
      context,
      namespace,
      resource_name,
      status
    )
SELECT
  uid AS resource_id,
  'Namespaces CPU default resource limit' AS title,
  context AS context,
  name AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(default_cpu_limit)
    FROM
      default_cpu_limits
    WHERE
      namespace = k8s_core_namespaces.name
      AND context = k8s_core_namespaces.context
  )
  = 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_core_namespaces;
```

### Namespaces CPU request resource quota

```sql
WITH
  default_request_cpu_limits
    AS (
      SELECT
        context,
        namespace,
        value->'default_request'->>'cpu' AS default_request_cpu_limit
      FROM
        k8s_core_limit_ranges
        CROSS JOIN jsonb_array_elements(k8s_core_limit_ranges.spec_limits)
    )
INSERT
INTO
  k8s_policy_results
    (
      resource_id,
      execution_time,
      framework,
      check_id,
      title,
      context,
      namespace,
      resource_name,
      status
    )
SELECT
  uid AS resource_id,
  'Namespaces CPU request resource quota' AS title,
  context AS context,
  name AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(default_request_cpu_limit)
    FROM
      default_request_cpu_limits
    WHERE
      namespace = k8s_core_namespaces.name
      AND context = k8s_core_namespaces.context
  )
  = 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_core_namespaces;
```

### Namespaces Memory default resource limit

```sql
WITH
  default_memory_limits
    AS (
      SELECT
        context, namespace, value->'default'->>'memory' AS default_memory_limit
      FROM
        k8s_core_limit_ranges
        CROSS JOIN jsonb_array_elements(k8s_core_limit_ranges.spec_limits)
    )
INSERT
INTO
  k8s_policy_results
    (
      resource_id,
      execution_time,
      framework,
      check_id,
      title,
      context,
      namespace,
      resource_name,
      status
    )
SELECT
  uid AS resource_id,
  'Namespaces Memory default resource limit' AS title,
  context AS context,
  name AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(default_memory_limit)
    FROM
      default_memory_limits
    WHERE
      namespace = k8s_core_namespaces.name
      AND context = k8s_core_namespaces.context
  )
  = 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_core_namespaces;
```

### Namespaces Memory request resource quota

```sql
WITH
  default_request_memory_limits
    AS (
      SELECT
        namespace,
        value->'default_request'->>'memory' AS default_request_memory_limit
      FROM
        k8s_core_limit_ranges
        CROSS JOIN jsonb_array_elements(k8s_core_limit_ranges.spec_limits)
    )
INSERT
INTO
  k8s_policy_results
    (
      resource_id,
      execution_time,
      framework,
      check_id,
      title,
      context,
      namespace,
      resource_name,
      status
    )
SELECT
  uid AS resource_id,
  'Namespaces Memory request resource quota' AS title,
  context AS context,
  name AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(default_request_memory_limit)
    FROM
      default_request_memory_limits
    WHERE
      namespace = k8s_core_namespaces.name
      AND context = k8s_core_namespaces.context
  )
  = 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_core_namespaces;
```


