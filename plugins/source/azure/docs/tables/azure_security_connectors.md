# Table: azure_security_connectors

This table shows data for Azure Security Connectors.

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity@v0.9.0#ConnectorSetting

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
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|