# Table: azure_healthbot_bots

This table shows data for Azure Healthbot Bots.

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot@v1.0.0#HealthBot

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|sku|`json`|
|identity|`json`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|