# Table: vault_sys_mounts

This table shows data for Vault Sys Mounts.

https://developer.hashicorp.com/vault/api-docs/system/mounts

The primary key for this table is **uuid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|path|`utf8`|
|uuid (PK)|`utf8`|
|type|`utf8`|
|description|`utf8`|
|accessor|`utf8`|
|config|`json`|
|options|`json`|
|local|`bool`|
|seal_wrap|`bool`|
|external_entropy_access|`bool`|
|plugin_version|`utf8`|
|running_plugin_version|`utf8`|
|running_sha256|`utf8`|
|deprecation_status|`utf8`|