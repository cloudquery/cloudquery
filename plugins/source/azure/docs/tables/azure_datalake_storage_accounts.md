
# Table: azure_datalake_storage_accounts
Data Lake Store account
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|identity_type|text|The type of encryption being used|
|identity_principal_id|uuid|The principal identifier associated with the encryption|
|identity_tenant_id|uuid|The tenant identifier associated with the encryption|
|default_group|text|The default owner group for all new folders and files created in the Data Lake Store account|
|encryption_config_type|text|The type of encryption configuration being used|
|encryption_config_key_vault_meta_info_key_vault_resource_id|text|The resource identifier for the user managed Key Vault being used to encrypt|
|encryption_config_key_vault_meta_info_encryption_key_name|text|The name of the user managed encryption key|
|encryption_config_key_vault_meta_info_encryption_key_version|text|The version of the user managed encryption key|
|encryption_state|text|The current state of encryption for this Data Lake Store account|
|encryption_provisioning_state|text|The current state of encryption provisioning for this Data Lake Store account|
|firewall_state|text|The current state of the IP address firewall for this Data Lake Store account|
|firewall_allow_azure_ips|text|The current state of allowing or disallowing IPs originating within Azure through the firewall|
|trusted_id_provider_state|text|The current state of the trusted identity provider feature for this Data Lake Store account|
|new_tier|text|The commitment tier to use for next month|
|current_tier|text|The commitment tier in use for the current month|
|account_id|uuid|The unique identifier associated with this Data Lake Store account|
|provisioning_state|text|The provisioning status of the Data Lake Store account|
|state|text|The state of the Data Lake Store account|
|creation_time|timestamp without time zone||
|last_modified_time|timestamp without time zone||
|endpoint|text|The full CName endpoint for this account|
|id|text|The resource identifier|
|name|text|The resource name|
|type|text|The resource type|
|location|text|The resource location|
|tags|jsonb|The resource tags|
