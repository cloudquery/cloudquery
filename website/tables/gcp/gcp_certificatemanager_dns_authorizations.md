# Table: gcp_certificatemanager_dns_authorizations

This table shows data for GCP Certificatemanager DNS Authorizations.

https://cloud.google.com/certificate-manager/docs/reference/rest/v1/projects.locations.dnsAuthorizations#DnsAuthorization

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|labels|`json`|
|description|`utf8`|
|domain|`utf8`|
|dns_resource_record|`json`|