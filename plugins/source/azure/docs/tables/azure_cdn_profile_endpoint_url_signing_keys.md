
# Table: azure_cdn_profile_endpoint_url_signing_keys
URLSigningKey url signing key
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|profile_endpoint_cq_id|uuid|Unique CloudQuery ID of azure_cdn_profile_endpoints table (FK)|
|key_id|text|Defines the customer defined key Id|
|key_source_parameters_odata_type|text||
|key_source_parameters_subscription_id|text|Subscription Id of the user's Key Vault containing the secret|
|key_source_parameters_resource_group_name|text|Resource group of the user's Key Vault containing the secret|
|key_source_parameters_vault_name|text|The name of the user's Key Vault containing the secret|
|key_source_parameters_secret_name|text|The name of secret in Key Vault|
|key_source_parameters_secret_version|text|The version(GUID) of secret in Key Vault|
