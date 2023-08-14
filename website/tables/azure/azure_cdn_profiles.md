# Table: azure_cdn_profiles

This table shows data for Azure Content Delivery Network (CDN) Profiles.

https://learn.microsoft.com/en-us/rest/api/cdn/profiles/list?tabs=HTTP#profile

The primary key for this table is **id**.

## Relations

The following tables depend on azure_cdn_profiles:
  - [azure_cdn_endpoints](azure_cdn_endpoints)
  - [azure_cdn_rule_sets](azure_cdn_rule_sets)
  - [azure_cdn_security_policies](azure_cdn_security_policies)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|location|`utf8`|
|sku|`json`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|kind|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|