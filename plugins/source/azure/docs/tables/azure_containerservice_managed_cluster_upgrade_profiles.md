# Table: azure_containerservice_managed_cluster_upgrade_profiles

https://learn.microsoft.com/en-us/rest/api/aks/managed-clusters/get-upgrade-profile?tabs=HTTP#managedclusterupgradeprofile

The primary key for this table is **id**.

## Relations

This table depends on [azure_containerservice_managed_clusters](azure_containerservice_managed_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|properties|JSON|
|id (PK)|String|
|name|String|
|type|String|