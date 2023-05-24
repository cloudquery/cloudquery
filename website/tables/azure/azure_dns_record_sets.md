# Table: azure_dns_record_sets

This table shows data for Azure DNS Record Sets.

https://learn.microsoft.com/en-us/rest/api/dns/record-sets/list-by-dns-zone?tabs=HTTP#recordset

The primary key for this table is **id**.

## Relations

This table depends on [azure_dns_zones](azure_dns_zones).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|etag|String|
|properties|JSON|
|id (PK)|String|
|name|String|
|type|String|