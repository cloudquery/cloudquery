# Table: k8s_core_resource_quotas

This table shows data for Kubernetes (K8s) Core Resource Quotas.

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
|spec_hard|`json`|
|spec_scopes|`list<item: utf8, nullable>`|
|spec_scope_selector|`json`|
|status_hard|`json`|
|status_used|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

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


