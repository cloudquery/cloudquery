# Table: k8s_networking_network_policies

This table shows data for Kubernetes (K8s) Networking Network Policies.

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
|spec_pod_selector|`json`|
|spec_ingress|`json`|
|spec_egress|`json`|
|spec_policy_types|`list<item: utf8, nullable>`|
|status_conditions|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

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


