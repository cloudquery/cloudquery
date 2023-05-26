# Table: gcp_dns_managed_zones

This table shows data for GCP DNS Managed Zones.

https://cloud.google.com/dns/docs/reference/v1/managedZones#resource

The primary key for this table is **id**.

## Relations

The following tables depend on gcp_dns_managed_zones:
  - [gcp_dns_resource_record_sets](gcp_dns_resource_record_sets)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|project_id|utf8|
|cloud_logging_config|json|
|creation_time|utf8|
|description|utf8|
|dns_name|utf8|
|dnssec_config|json|
|forwarding_config|json|
|id (PK)|int64|
|kind|utf8|
|labels|json|
|name|utf8|
|name_server_set|utf8|
|name_servers|list<item: utf8, nullable>|
|peering_config|json|
|private_visibility_config|json|
|reverse_lookup_config|json|
|service_directory_config|json|
|visibility|utf8|