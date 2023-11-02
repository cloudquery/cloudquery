# Table: k8s_core_service_accounts

This table shows data for Kubernetes (K8s) Core Service Accounts.

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
|secrets|`json`|
|image_pull_secrets|`json`|
|automount_service_account_token|`bool`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

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
  WHEN automount_service_account_token THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  k8s_core_service_accounts;
```


