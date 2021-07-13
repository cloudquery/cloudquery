
# Table: azure_keyvault_vault_secrets
Azure keyvault secrets
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|vault_cq_id|uuid|Unique ID of azure_keyvault_vaults table (FK)|
|id|text|Secret identifier|
|recoverable_days|integer|softDelete data retention days Value should be >=7 and <=90 when softDelete enabled, otherwise 0|
|recovery_level|text|Reflects the deletion recovery level currently in effect for secrets in the current vault If it contains 'Purgeable', the secret can be permanently deleted by a privileged user; otherwise, only the system can purge the secret, at the end of the retention interval Possible values include: 'Purgeable', 'RecoverablePurgeable', 'Recoverable', 'RecoverableProtectedSubscription', 'CustomizedRecoverablePurgeable', 'CustomizedRecoverable', 'CustomizedRecoverableProtectedSubscription'|
|enabled|boolean|Determines whether the object is enabled|
|tags|jsonb|Application specific metadata in the form of key-value pairs|
|content_type|text|Type of the secret value such as a password|
|managed|boolean|True if the secret's lifetime is managed by key vault If this is a key backing a certificate, then managed will be true|
