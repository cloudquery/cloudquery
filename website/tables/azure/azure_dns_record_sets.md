# Table: azure_dns_record_sets

This table shows data for Azure DNS Record Sets.

https://learn.microsoft.com/en-us/rest/api/dns/record-sets/list-by-dns-zone?tabs=HTTP#recordset

The primary key for this table is **id**.

## Relations

This table depends on [azure_dns_zones](azure_dns_zones).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|etag|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|