# Table: awspricing_service_products

This table shows data for Awspricing Service Products.

The primary key for this table is **sku**.

## Relations

This table depends on [awspricing_services](awspricing_services).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|sku (PK)|String|
|product_family|String|
|attributes|JSON|