# Table: azure_apimanagement_service

This table shows data for Azure API Management Service.

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement#ServiceResource

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
|zones|`list<item: utf8, nullable>`|
|etag|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|