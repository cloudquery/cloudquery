# Table: gcp_logging_metrics

This table shows data for GCP Logging Metrics.

https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.metrics#LogMetric

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|description|`utf8`|
|filter|`utf8`|
|disabled|`bool`|
|metric_descriptor|`json`|
|value_extractor|`utf8`|
|label_extractors|`json`|
|bucket_options|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|version|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that the log metric filter and alerts exist for Audit Configuration changes (Automated)

```sql
SELECT
  filter AS resource_id,
  'Ensure that the log metric filter and alerts exist for Audit Configuration changes (Automated)'
    AS title,
  project_id AS project_id,
  CASE
  WHEN disabled = false
  AND filter
    ~ e'\\s*protoPayload.methodName\\s*=\\s*"SetIamPolicy"\\s*AND\\s*protoPayload.serviceData.policyDelta.auditConfigDeltas:*\\s*'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_logging_metrics;
```

### Ensure that the log metric filter and alerts exist for Custom Role changes (Automated)

```sql
SELECT
  filter AS resource_id,
  'Ensure that the log metric filter and alerts exist for Custom Role changes (Automated)'
    AS title,
  project_id AS project_id,
  CASE
  WHEN disabled = false
  AND filter
    ~ e'\\s*resource.type\\s*=\\s*"iam_role"\\s*AND\\s*protoPayload.methodName\\s*=\\s*"google.iam.admin.v1.CreateRole"\\s*OR\\s*protoPayload.methodName\\s*=\\s*"google.iam.admin.v1.DeleteRole"\\s*OR\\s*protoPayload.methodName\\s*=\\s*"google.iam.admin.v1.UpdateRole"\\s*'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_logging_metrics;
```

### Ensure log metric filter and alerts exist for project ownership assignments/changes (Automated)

```sql
SELECT
  filter AS resource_id,
  'Ensure log metric filter and alerts exist for project ownership assignments/changes (Automated)'
    AS title,
  project_id AS project_id,
  CASE
  WHEN disabled = false
  AND filter
    ~ e'\\s*(\\s*protoPayload.serviceName\\s*=\\s*"cloudresourcemanager.googleapis.com"\\s*)\\s*AND\\s*(\\s*ProjectOwnership\\s*OR\\s*projectOwnerInvitee\\s*)\\s*OR\\s*(\\s*protoPayload.serviceData.policyDelta.bindingDeltas.action\\s*=\\s*"REMOVE"\\s*AND\\s*protoPayload.serviceData.policyDelta.bindingDeltas.role\\s*=\\s*"roles/owner"\\s*)\\s*OR\\s*(\\s*protoPayload.serviceData.policyDelta.bindingDeltas.action\\s*=\\s*"ADD"\\s*AND\\s*protoPayload.serviceData.policyDelta.bindingDeltas.role\\s*=\\s*"roles/owner"\\s*)\\s*'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_logging_metrics;
```

### Ensure that the log metric filter and alerts exist for SQL instance configuration changes (Automated)

```sql
SELECT
  filter AS resource_id,
  'Ensure that the log metric filter and alerts exist for SQL instance configuration changes (Automated)'
    AS title,
  project_id AS project_id,
  CASE
  WHEN disabled = false
  AND filter = 'protoPayload.methodName="cloudsql.instances.update"'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_logging_metrics;
```

### Ensure that the log metric filter and alerts exist for Cloud Storage IAM permission changes (Automated)

```sql
SELECT
  filter AS resource_id,
  'Ensure that the log metric filter and alerts exist for Cloud Storage IAM permission changes (Automated)'
    AS title,
  project_id AS project_id,
  CASE
  WHEN disabled = false
  AND filter
    ~ e'\\s*resource.type\\s*=\\s*gcs_bucket\\s*AND\\s*protoPayload.methodName\\s*=\\s*"storage.setIamPermissions"\\s*'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_logging_metrics;
```

### Ensure that the log metric filter and alerts exist for VPC Network Firewall rule changes (Automated)

```sql
SELECT
  filter AS resource_id,
  'Ensure that the log metric filter and alerts exist for VPC Network Firewall rule changes (Automated)'
    AS title,
  project_id AS project_id,
  CASE
  WHEN disabled = false
  AND filter
    ~ e'\\s*resource.type\\s*=\\s*"gce_firewall_rule"\\s*AND\\s*protoPayload.methodName\\s*=\\s*"v1.compute.firewalls.patch"\\s*OR\\s*protoPayload.methodName\\s*=\\s*"v1.compute.firewalls.insert"\\s*'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_logging_metrics;
```

### Ensure that the log metric filter and alerts exist for VPC network changes (Automated)

```sql
SELECT
  filter AS resource_id,
  'Ensure that the log metric filter and alerts exist for VPC network changes (Automated)'
    AS title,
  project_id AS project_id,
  CASE
  WHEN disabled = false
  AND filter
    ~ e'\\s*resource.type\\s*=\\s*gce_network\\s*AND\\s*protoPayload.methodName\\s*=\\s*"beta.compute.networks.insert"\\s*OR\\s*protoPayload.methodName\\s*=\\s*"beta.compute.networks.patch"\\s*OR\\s*protoPayload.methodName\\s*=\\s*"v1.compute.networks.delete"\\s*OR\\s*protoPayload.methodName\\s*=\\s*"v1.compute.networks.removePeering"\\s*OR\\s*protoPayload.methodName\\s*=\\s*"v1.compute.networks.addPeering"\\s*'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_logging_metrics;
```

### Ensure that the log metric filter and alerts exist for VPC network route changes (Automated)

```sql
SELECT
  filter AS resource_id,
  'Ensure that the log metric filter and alerts exist for VPC network route changes (Automated)'
    AS title,
  project_id AS project_id,
  CASE
  WHEN disabled = false
  AND filter
    ~ e'\\s*resource.type\\s*=\\s*"gce_route"\\s*AND\\s*protoPayload.methodName\\s*=\\s*"beta.compute.routes.patch"\\s*OR\\s*protoPayload.methodName\\s*=\\s*"beta.compute.routes.insert"\\s*'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_logging_metrics;
```


