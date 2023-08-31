# Table: awspricing_service_products

This table shows data for Service Products from the AWS Price List API.

https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/price-changes.html

The primary key for this table is **sku**.

## Relations

This table depends on [awspricing_services](awspricing_services).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|sku (PK)|`utf8`|
|product_family|`utf8`|
|attributes|`json`|