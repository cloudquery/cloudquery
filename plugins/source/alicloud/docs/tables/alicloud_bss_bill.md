# Table: alicloud_bss_bill

https://help.aliyun.com/document_detail/100400.html

The composite primary key for this table is (**billing_cycle**, **account_id**, **subscription_type**, **commodity_code**, **product_code**, **product_type**, **pip_code**, **record_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|billing_cycle (PK)|String|
|account_id (PK)|String|
|account_name|String|
|product_name|String|
|sub_order_id|String|
|deducted_by_cash_coupons|Float|
|payment_time|String|
|payment_amount|Float|
|deducted_by_prepaid_card|Float|
|invoice_discount|Float|
|usage_end_time|String|
|item|String|
|subscription_type (PK)|String|
|pretax_gross_amount|Float|
|currency|String|
|commodity_code (PK)|String|
|usage_start_time|String|
|adjust_amount|Float|
|status|String|
|deducted_by_coupons|Float|
|round_down_discount|String|
|product_detail|String|
|product_code (PK)|String|
|product_type (PK)|String|
|outstanding_amount|Float|
|pip_code (PK)|String|
|pretax_amount|Float|
|owner_id|String|
|record_id (PK)|String|
|cash_amount|Float|