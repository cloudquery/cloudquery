# Table: vault_sys_plugins

This table shows data for Vault Sys Plugins.

The composite primary key for this table is (**type**, **name**, **version**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|type (PK)|`utf8`|
|name (PK)|`utf8`|
|version (PK)|`utf8`|
|builtin|`bool`|
|deprecation_status|`utf8`|