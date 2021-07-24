
# Table: azure_keyvault_vault_keys
KeyItem the key item containing key metadata
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|vault_cq_id|uuid|Unique CloudQuery ID of azure_keyvault_vaults table (FK)|
|kid|text|Key identifier|
|recoverable_days|integer|softDelete data retention days Value should be >=7 and <=90 when softDelete enabled, otherwise 0|
|recovery_level|text|Reflects the deletion recovery level currently in effect for keys in the current vault If it contains 'Purgeable' the key can be permanently deleted by a privileged user; otherwise, only the system can purge the key, at the end of the retention interval Possible values include: 'Purgeable', 'RecoverablePurgeable', 'Recoverable', 'RecoverableProtectedSubscription', 'CustomizedRecoverablePurgeable', 'CustomizedRecoverable', 'CustomizedRecoverableProtectedSubscription'|
|enabled|boolean|Determines whether the object is enabled|
|not_before|timestamp without time zone|Not before date in UTC|
|expires|timestamp without time zone|Expiry date in UTC|
|created|timestamp without time zone|Creation time in UTC|
|updated|timestamp without time zone|Last updated time in UTC|
|tags|jsonb|Application specific metadata in the form of key-value pairs|
|managed|boolean|True if the key's lifetime is managed by key vault If this is a key backing a certificate, then managed will be true|
