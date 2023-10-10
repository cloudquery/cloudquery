# Table: vault_sys_audits

This table shows data for Vault Sys Audits.

https://developer.hashicorp.com/vault/api-docs/system/audit

The composite primary key for this table is (**type**, **path**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|type (PK)|`utf8`|
|description|`utf8`|
|options|`json`|
|local|`bool`|
|path (PK)|`utf8`|