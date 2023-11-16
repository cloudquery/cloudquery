# Table: azure_saas_resources

This table shows data for Azure Software as a Service (SaaS) Resources.

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas@v0.5.0#Resource

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|