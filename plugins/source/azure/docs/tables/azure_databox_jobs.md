# Table: azure_databox_jobs

This table shows data for Azure Data Box Jobs.

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox@v1.0.0#JobResource

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|properties|`json`|
|sku|`json`|
|identity|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|