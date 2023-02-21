# Table: azure_cdn_profiles

https://learn.microsoft.com/en-us/rest/api/cdn/profiles/list?tabs=HTTP#profile

The primary key for this table is **id**.

## Relations

The following tables depend on azure_cdn_profiles:
  - [azure_cdn_endpoints](azure_cdn_endpoints.md)
  - [azure_cdn_rule_sets](azure_cdn_rule_sets.md)
  - [azure_cdn_security_policies](azure_cdn_security_policies.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|location|String|
|sku|JSON|
|properties|JSON|
|tags|JSON|
|id (PK)|String|
|kind|String|
|name|String|
|system_data|JSON|
|type|String|