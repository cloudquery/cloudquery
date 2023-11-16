# Table: azure_operationalinsights_clusters

This table shows data for Azure Operational Insights Clusters.

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights@v1.0.0#Cluster

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|identity|`json`|
|properties|`json`|
|sku|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|