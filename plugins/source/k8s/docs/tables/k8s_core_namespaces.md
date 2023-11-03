# Table: k8s_core_namespaces

This table shows data for Kubernetes (K8s) Core Namespaces.

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
|spec_finalizers|`list<item: utf8, nullable>`|
|status_phase|`utf8`|
|status_conditions|`json`|

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

### Namespace enforces resource quota cpu limits

```sql
SELECT
  DISTINCT
  k8s_core_namespaces.uid AS resource_id,
  'Namespace enforces resource quota cpu limits' AS title,
  k8s_core_namespaces.context AS context,
  k8s_core_namespaces.name AS namespace,
  k8s_core_namespaces.name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      k8s_core_resource_quotas
    WHERE
      namespace = k8s_core_namespaces.name
      AND context = k8s_core_namespaces.context
      AND (spec_hard->>'limits.cpu') IS NOT NULL
  )
  = 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_core_namespaces;
```

### Namespace enforces resource quota cpu request

```sql
SELECT
  DISTINCT
  k8s_core_namespaces.uid AS resource_id,
  'Namespace enforces resource quota cpu request' AS title,
  k8s_core_namespaces.context AS context,
  k8s_core_namespaces.name AS namespace,
  k8s_core_namespaces.name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      k8s_core_resource_quotas
    WHERE
      namespace = k8s_core_namespaces.name
      AND context = k8s_core_namespaces.context
      AND (spec_hard->>'requests.cpu') IS NOT NULL
  )
  = 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_core_namespaces;
```

### Namespace enforces resource quota memory limits

```sql
SELECT
  DISTINCT
  k8s_core_namespaces.uid AS resource_id,
  'Namespace enforces resource quota memory limits' AS title,
  k8s_core_namespaces.context AS context,
  k8s_core_namespaces.name AS namespace,
  k8s_core_namespaces.name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      k8s_core_resource_quotas
    WHERE
      namespace = k8s_core_namespaces.name
      AND context = k8s_core_namespaces.context
      AND (spec_hard->>'limits.memory') IS NOT NULL
  )
  = 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_core_namespaces;
```

### Namespace enforces resource quota memory request

```sql
SELECT
  DISTINCT
  k8s_core_namespaces.uid AS resource_id,
  'Namespace enforces resource quota memory request' AS title,
  k8s_core_namespaces.context AS context,
  k8s_core_namespaces.name AS namespace,
  k8s_core_namespaces.name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      k8s_core_resource_quotas
    WHERE
      namespace = k8s_core_namespaces.name
      AND context = k8s_core_namespaces.context
      AND (spec_hard->>'requests.memory') IS NOT NULL
  )
  = 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_core_namespaces;
```

### Network policy default deny egress

```sql
SELECT
  uid AS resource_id,
  'Network policy default deny egress' AS title,
  context AS context,
  name AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      k8s_networking_network_policies
    WHERE
      namespace = k8s_core_namespaces.name
      AND context = k8s_core_namespaces.context
      AND spec_policy_types @> ARRAY['Egress']
      AND spec_pod_selector::STRING = '{}'
      AND ((spec_egress IS NULL) OR jsonb_array_length(spec_egress) = 0)
  )
  = 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_core_namespaces;
```

### Network policy default deny ingress

```sql
SELECT
  uid AS resource_id,
  'Network policy default deny ingress' AS title,
  context AS context,
  name AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      k8s_networking_network_policies
    WHERE
      namespace = k8s_core_namespaces.name
      AND context = k8s_core_namespaces.context
      AND spec_policy_types @> ARRAY['Ingress']
      AND spec_pod_selector::STRING = '{}'
      AND ((spec_ingress IS NULL) OR jsonb_array_length(spec_ingress) = 0)
  )
  = 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_core_namespaces;
```


