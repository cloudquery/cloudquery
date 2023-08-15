# Table: awspricing_service_terms

This table shows data for Service Terms from the AWS Price List API.

https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/price-changes.html

The composite primary key for this table is (**offer_term_code**, **sku**).

## Relations

This table depends on [awspricing_services](awspricing_services).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|type|`utf8`|
|offer_term_code (PK)|`utf8`|
|sku (PK)|`utf8`|
|effective_date|`timestamp[us, tz=UTC]`|
|price_dimensions|`json`|
|term_attributes|`json`|