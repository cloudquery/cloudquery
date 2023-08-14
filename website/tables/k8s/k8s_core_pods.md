# Table: k8s_core_pods

This table shows data for Kubernetes (K8s) Core Pods.

The primary key for this table is **uid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|context|`utf8`|
|status_host_ip|`inet`|
|status_pod_ip|`inet`|
|status_pod_ips|`list<item: inet, nullable>`|
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
|spec_volumes|`json`|
|spec_init_containers|`json`|
|spec_containers|`json`|
|spec_ephemeral_containers|`json`|
|spec_restart_policy|`utf8`|
|spec_termination_grace_period_seconds|`int64`|
|spec_active_deadline_seconds|`int64`|
|spec_dns_policy|`utf8`|
|spec_node_selector|`json`|
|spec_service_account_name|`utf8`|
|spec_automount_service_account_token|`bool`|
|spec_node_name|`utf8`|
|spec_host_network|`bool`|
|spec_host_pid|`bool`|
|spec_host_ipc|`bool`|
|spec_share_process_namespace|`bool`|
|spec_security_context|`json`|
|spec_image_pull_secrets|`json`|
|spec_hostname|`utf8`|
|spec_subdomain|`utf8`|
|spec_affinity|`json`|
|spec_scheduler_name|`utf8`|
|spec_tolerations|`json`|
|spec_host_aliases|`json`|
|spec_priority_class_name|`utf8`|
|spec_priority|`int64`|
|spec_dns_config|`json`|
|spec_readiness_gates|`json`|
|spec_runtime_class_name|`utf8`|
|spec_enable_service_links|`bool`|
|spec_preemption_policy|`utf8`|
|spec_overhead|`json`|
|spec_topology_spread_constraints|`json`|
|spec_set_hostname_as_fqdn|`bool`|
|spec_os|`json`|
|spec_host_users|`bool`|
|spec_scheduling_gates|`json`|
|spec_resource_claims|`json`|
|status_phase|`utf8`|
|status_conditions|`json`|
|status_message|`utf8`|
|status_reason|`utf8`|
|status_nominated_node_name|`utf8`|
|status_start_time|`timestamp[us, tz=UTC]`|
|status_init_container_statuses|`json`|
|status_container_statuses|`json`|
|status_qos_class|`utf8`|
|status_ephemeral_container_statuses|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Pod container privileged access disabled

```sql
WITH
  pod_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_core_pods CROSS JOIN jsonb_array_elements(spec_containers) AS value
    )
SELECT
  uid AS resource_id,
  'Pod container privileged access disabled' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      pod_containers
    WHERE
      pod_containers.uid = k8s_core_pods.uid
      AND pod_containers.container->'securityContext'->>'privileged' = 'true'
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_core_pods;
```

### Pod container privilege escalation disabled

```sql
WITH
  pod_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_core_pods CROSS JOIN jsonb_array_elements(spec_containers) AS value
    )
SELECT
  uid AS resource_id,
  'Pod container privilege escalation disabled' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      pod_containers
    WHERE
      pod_containers.uid = k8s_core_pods.uid
      AND pod_containers.container->'securityContext'->>'allowPrivilegeEscalation'
        = 'true'
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_core_pods;
```

### Pods container hostNetwork disabled

```sql
SELECT
  uid AS resource_id,
  'Pods container hostNetwork disabled' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE WHEN spec_host_network THEN 'fail' ELSE 'pass' END AS status
FROM
  k8s_core_pods;
```

### Pod containers HostPID and HostIPC sharing disabled

```sql
SELECT
  uid AS resource_id,
  'Pod containers HostPID and HostIPC sharing disabled' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE WHEN spec_host_pid OR spec_host_ipc THEN 'fail' ELSE 'pass' END AS status
FROM
  k8s_core_pods;
```

### Pod container filesystem is read-ony

```sql
WITH
  pod_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_core_pods CROSS JOIN jsonb_array_elements(spec_containers) AS value
    )
SELECT
  uid AS resource_id,
  'Pod container filesystem is read-ony' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      pod_containers
    WHERE
      pod_containers.uid = k8s_core_pods.uid
      AND pod_containers.container->'securityContext'->>'readOnlyRootFilesystem'
        IS DISTINCT FROM 'true'
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_core_pods;
```

### Pod container runs as non-root

```sql
WITH
  pod_containers
    AS (
      SELECT
        uid, value AS container
      FROM
        k8s_core_pods CROSS JOIN jsonb_array_elements(spec_containers) AS value
    )
SELECT
  uid AS resource_id,
  'Pod container runs as non-root' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      pod_containers
    WHERE
      pod_containers.uid = k8s_core_pods.uid
      AND pod_containers.container->'securityContext'->>'runAsNonRoot'
        IS DISTINCT FROM 'true'
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_core_pods;
```

### Pod service account tokens disabled

```sql
SELECT
  DISTINCT
  uid AS resource_id,
  'Pod service account tokens disabled' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN spec_automount_service_account_token THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_core_pods;
```

### Pod volume don''t have a hostPath

```sql
WITH
  pod_volumes
    AS (
      SELECT
        uid, value AS volumes
      FROM
        k8s_core_pods CROSS JOIN jsonb_array_elements(spec_volumes) AS value
    )
SELECT
  uid AS resource_id,
  e'Pod volume don\'t have a hostPath' AS title,
  context AS context,
  namespace AS namespace,
  name AS resource_name,
  CASE
  WHEN (
    SELECT
      count(*)
    FROM
      pod_volumes
    WHERE
      pod_volumes.uid = k8s_core_pods.uid
      AND (pod_volumes.volumes->>'hostPath') IS NOT NULL
  )
  > 0
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_core_pods;
```


