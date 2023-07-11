# Table: azure_storage_file_shares

This table shows data for Azure Storage File Shares.

https://learn.microsoft.com/en-us/rest/api/storagerp/file-shares/list?tabs=HTTP#fileshareitem

The primary key for this table is **id**.

## Relations

This table depends on [azure_storage_accounts](azure_storage_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|etag|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|