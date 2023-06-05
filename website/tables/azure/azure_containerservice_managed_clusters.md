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
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
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