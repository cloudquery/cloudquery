# Table: azure_compute_ssh_public_keys

This table shows data for Azure Compute SSH Public Keys.

https://learn.microsoft.com/en-us/rest/api/compute/ssh-public-keys/list-by-subscription?tabs=HTTP#sshpublickeyresource

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|