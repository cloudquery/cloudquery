# Table: azure_compute_disk_encryption_sets

This table shows data for Azure Compute Disk Encryption Sets.

https://learn.microsoft.com/en-us/rest/api/compute/disk-encryption-sets/list?tabs=HTTP#diskencryptionset

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|identity|`json`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|