# Table: azure_botservice_bots

This table shows data for Azure Botservice Bots.

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice@v0.5.0#Bot

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|etag|`utf8`|
|kind|`utf8`|
|location|`utf8`|
|properties|`json`|
|sku|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|
|zones|`list<item: utf8, nullable>`|