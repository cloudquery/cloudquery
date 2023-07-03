# Table: gcp_dns_resource_record_sets

This table shows data for GCP DNS Resource Record Sets.

https://cloud.google.com/dns/docs/reference/v1/resourceRecordSets

The composite primary key for this table is (**project_id**, **managed_zone_name**, **name**, **type**).

## Relations

This table depends on [gcp_dns_managed_zones](gcp_dns_managed_zones).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|managed_zone_name (PK)|`utf8`|
|kind|`utf8`|
|name (PK)|`utf8`|
|routing_policy|`json`|
|rrdatas|`list<item: utf8, nullable>`|
|signature_rrdatas|`list<item: utf8, nullable>`|
|ttl|`int64`|
|type (PK)|`utf8`|