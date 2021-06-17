
# Table: azure_keyvault_vault_keys
Azure ketvault vault key
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|vault_id|uuid|Unique ID of azure_keyvault_vaults table (FK)|
|attributes_enabled|boolean|Determines whether or not the object is enabled|
|attributes_not_before|bigint|Not before date in seconds since 1970-01-01T00:00:00Z|
|attributes_expires|bigint|Expiry date in seconds since 1970-01-01T00:00:00Z|
|attributes_created|bigint|Creation time in seconds since 1970-01-01T00:00:00Z|
|attributes_updated|bigint|Last updated time in seconds since 1970-01-01T00:00:00Z|
|attributes_recovery_level|text|The deletion recovery level currently in effect for the object If it contains 'Purgeable', then the object can be permanently deleted by a privileged user; otherwise, only the system can purge the object at the end of the retention interval Possible values include: 'Purgeable', 'RecoverablePurgeable', 'Recoverable', 'RecoverableProtectedSubscription'|
|kty|text|The type of the key For valid values, see JsonWebKeyType Possible values include: 'EC', 'ECHSM', 'RSA', 'RSAHSM'|
|key_ops|text[]|Enumerates the values for json web key operation|
|key_size|integer|The key size in bits For example: 2048, 3072, or 4096 for RSA|
|curve_name|text|The elliptic curve name For valid values, see JsonWebKeyCurveName Possible values include: 'P256', 'P384', 'P521', 'P256K'|
|key_uri|text|The URI to retrieve the current version of the key|
|key_uri_with_version|text|The URI to retrieve the specific version of the key|
|resource_id|text|Fully qualified identifier of the key vault resource|
|name|text|Name of the key vault resource|
|type|text|Resource type of the key vault resource|
|location|text|Azure location of the key vault resource|
|tags|jsonb|Tags assigned to the key vault resource|
