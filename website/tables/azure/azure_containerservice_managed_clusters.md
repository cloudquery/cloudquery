# Table: azure_containerservice_managed_clusters

This table shows data for Azure Container Service Managed Clusters.

https://learn.microsoft.com/en-us/rest/api/aks/managed-clusters/list?tabs=HTTP#managedcluster

The primary key for this table is **id**.

## Relations

The following tables depend on azure_containerservice_managed_clusters:
  - [azure_containerservice_managed_cluster_upgrade_profiles](azure_containerservice_managed_cluster_upgrade_profiles)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|extended_location|`json`|
|identity|`json`|
|properties|`json`|
|sku|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### External accounts with owner permissions should be removed from your subscription

```sql
SELECT
  'External accounts with owner permissions should be removed from your subscription'
    AS title,
  mc.subscription_id AS subscription_id,
  mc.id AS resource_id,
  CASE
  WHEN (properties->>'enableRBAC')::BOOL IS NOT true THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_containerservice_managed_clusters AS mc
  INNER JOIN azure_subscription_subscriptions AS sub ON
      sub.id = mc.subscription_id;
```

### Role-Based Access Control (RBAC) should be used on Kubernetes Services

```sql
SELECT
  'Role-Based Access Control (RBAC) should be used on Kubernetes Services'
    AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN (properties->>'enableRBAC')::BOOL IS NOT true THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_containerservice_managed_clusters;
```


