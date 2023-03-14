# Table: awspricing_services

This table shows data for Services from the AWS Price List API.

The composite primary key for this table is (**offer_code**, **version**, **publication_date**).

## Relations

The following tables depend on awspricing_services:
  - [awspricing_service_products](awspricing_service_products)
  - [awspricing_service_terms](awspricing_service_terms)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|format_version|String|
|disclaimer|String|
|offer_code (PK)|String|
|version (PK)|String|
|publication_date (PK)|Timestamp|