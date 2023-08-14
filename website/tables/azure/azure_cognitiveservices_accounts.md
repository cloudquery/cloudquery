# Table: azure_cognitiveservices_accounts

This table shows data for Azure Cognitive Services Accounts.

https://learn.microsoft.com/en-us/rest/api/cognitiveservices/accountmanagement/accounts/list?tabs=HTTP#account

The primary key for this table is **id**.

## Relations

The following tables depend on azure_cognitiveservices_accounts:
  - [azure_cognitiveservices_account_deployments](azure_cognitiveservices_account_deployments)
  - [azure_cognitiveservices_account_models](azure_cognitiveservices_account_models)
  - [azure_cognitiveservices_account_private_endpoint_connections](azure_cognitiveservices_account_private_endpoint_connections)
  - [azure_cognitiveservices_account_private_link_resources](azure_cognitiveservices_account_private_link_resources)
  - [azure_cognitiveservices_account_skus](azure_cognitiveservices_account_skus)
  - [azure_cognitiveservices_account_usages](azure_cognitiveservices_account_usages)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|identity|`json`|
|kind|`utf8`|
|location|`utf8`|
|properties|`json`|
|sku|`json`|
|tags|`json`|
|etag|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|