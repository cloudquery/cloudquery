# Table: azure_cognitiveservices_account_models

This table shows data for Azure Cognitive Services Account Models.

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices@v1.3.0#AccountModel

The primary key for this table is **_cq_id**.

## Relations

This table depends on [azure_cognitiveservices_accounts](azure_cognitiveservices_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|account_id|`utf8`|
|base_model|`json`|
|capabilities|`json`|
|deprecation|`json`|
|finetune_capabilities|`json`|
|format|`utf8`|
|lifecycle_status|`utf8`|
|max_capacity|`int64`|
|name|`utf8`|
|version|`utf8`|
|call_rate_limit|`json`|
|system_data|`json`|