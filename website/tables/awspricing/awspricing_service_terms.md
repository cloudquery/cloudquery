# Table: awspricing_service_terms

This table shows data for Service Terms from the AWS Price List API.

https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/price-changes.html

The composite primary key for this table is (**offer_term_code**, **sku**).

## Relations

This table depends on [awspricing_services](awspricing_services).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|type|String|
|offer_term_code (PK)|String|
|sku (PK)|String|
|effective_date|Timestamp|
|price_dimensions|JSON|
|term_attributes|JSON|