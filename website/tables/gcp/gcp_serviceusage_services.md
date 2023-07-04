# Table: gcp_serviceusage_services

This table shows data for GCP Serviceusage Services.

https://cloud.google.com/service-usage/docs/reference/rest/v1/services#Service

The primary key for this table is **name**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|name (PK)|`utf8`|
|parent|`utf8`|
|config|`json`|
|state|`utf8`|