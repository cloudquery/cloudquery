# Table: googleads_customer_labels

This table shows data for Google Ads Customer Labels.

https://developers.google.com/google-ads/api/reference/rpc/v13/CustomerLabel

The composite primary key for this table is (**customer_id**, **resource_name**, **customer**).

## Relations

This table depends on [googleads_customers](googleads_customers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|customer_id (PK)|Int|
|resource_name (PK)|String|
|customer (PK)|String|
|label|String|