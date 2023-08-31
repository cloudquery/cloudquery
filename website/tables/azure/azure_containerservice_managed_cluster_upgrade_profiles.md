# Table: azure_containerservice_managed_cluster_upgrade_profiles

This table shows data for Azure Container Service Managed Cluster Upgrade Profiles.

https://learn.microsoft.com/en-us/rest/api/aks/managed-clusters/get-upgrade-profile?tabs=HTTP#managedclusterupgradeprofile

The primary key for this table is **id**.

## Relations

This table depends on [azure_containerservice_managed_clusters](azure_containerservice_managed_clusters).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|