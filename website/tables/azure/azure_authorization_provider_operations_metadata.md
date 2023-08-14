# Table: azure_authorization_provider_operations_metadata

This table shows data for Azure Authorization Provider Operations Metadata.

https://learn.microsoft.com/en-us/rest/api/authorization/provider-operations-metadata/list?tabs=HTTP#provideroperationsmetadata

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|display_name|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|operations|`json`|
|resource_types|`json`|
|type|`utf8`|