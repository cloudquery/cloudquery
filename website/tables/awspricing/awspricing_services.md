# Table: awspricing_services

This table shows data for Services from the AWS Price List API.

https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/price-changes.html

The composite primary key for this table is (**offer_code**, **version**, **publication_date**).

## Relations

The following tables depend on awspricing_services:
  - [awspricing_service_products](awspricing_service_products)
  - [awspricing_service_terms](awspricing_service_terms)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|format_version|`utf8`|
|disclaimer|`utf8`|
|offer_code (PK)|`utf8`|
|version (PK)|`utf8`|
|publication_date (PK)|`timestamp[us, tz=UTC]`|