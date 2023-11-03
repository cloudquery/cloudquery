# Table: azure_compute_availability_sets

This table shows data for Azure Compute Availability Sets.

https://learn.microsoft.com/en-us/rest/api/compute/availability-sets/list-by-subscription?tabs=HTTP#availabilityset

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|properties|`json`|
|sku|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|